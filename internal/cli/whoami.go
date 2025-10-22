package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/helixprojectai-code/helix-sdk-go/pkg/auth"
)

var (
	whoamiBaseURL string
	whoamiTimeout time.Duration
)

var whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Show the authenticated identity using REAL JWT tokens",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Load cached token
		token, err := auth.LoadToken()
		if err != nil {
			return fmt.Errorf("no cached token found: %w", err)
		}

		fmt.Printf("🔑 Using cached JWT token (length: %d)...\n", len(token))

		// Create HTTP request directly
		ctx, cancel := context.WithTimeout(context.Background(), whoamiTimeout)
		defer cancel()

		req, err := http.NewRequestWithContext(ctx, "GET", whoamiBaseURL+"/meta/whoami", nil)
		if err != nil {
			return fmt.Errorf("create request: %w", err)
		}

		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("User-Agent", "helix-bridgectl/1.0")

		client := &http.Client{Timeout: whoamiTimeout}
		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("execute request: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			return fmt.Errorf("whoami failed: %s - %s", resp.Status, string(body))
		}

		// Parse response
		var identity struct {
			Subject      string   `json:"subject"`
			Tenant       string   `json:"tenant"`
			Scopes       []string `json:"scopes"`
			TokenPreview string   `json:"token_preview"`
			OK           bool     `json:"ok"`
			Exp          interface{} `json:"exp,omitempty"`
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("read response: %w", err)
		}

		if err := json.Unmarshal(body, &identity); err != nil {
			return fmt.Errorf("decode response: %w", err)
		}

		// Pretty output with JWT details
		fmt.Printf("🔐 Authenticated Identity (REAL JWT):\n")
		fmt.Printf("   Subject: %s\n", identity.Subject)
		fmt.Printf("   Tenant:  %s\n", identity.Tenant)
		fmt.Printf("   Scopes:  %s\n", strings.Join(identity.Scopes, ", "))
		fmt.Printf("   Token:   %s\n", identity.TokenPreview)
		if identity.Exp != nil {
			expTime := time.Unix(int64(identity.Exp.(float64)), 0)
			fmt.Printf("   Expires: %s (%s)\n", expTime.Format("2006-01-02 15:04:05"), time.Until(expTime).Round(time.Minute))
		}
		fmt.Printf("   Status:  ✅ Authenticated with REAL JWT\n")

		return nil
	},
}

func init() {
	whoamiCmd.Flags().StringVar(&whoamiBaseURL, "base-url", "http://127.0.0.1:3000", "Bridge API base URL")
	whoamiCmd.Flags().DurationVar(&whoamiTimeout, "timeout", 5*time.Second, "Request timeout")
}
