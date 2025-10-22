package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/helixprojectai-code/helix-sdk-go/pkg/auth"
	"github.com/helixprojectai-code/helix-sdk-go/pkg/client"
	"github.com/helixprojectai-code/helix-sdk-go/pkg/config"
)

var (
	loginBaseURL    string
	loginClientID   string
	loginSecret     string
	loginTenantID   string
	loginTimeout    time.Duration
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate with the Helix Bridge API and cache the JWT token",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := config.DefaultConfig().
			WithBaseURL(loginBaseURL).
			WithTimeout(loginTimeout)

		c := client.New(cfg)

		ctx, cancel := context.WithTimeout(context.Background(), loginTimeout)
		defer cancel()

		// Test connectivity first
		if err := c.Ping(ctx); err != nil {
			return fmt.Errorf("cannot reach bridge API: %w", err)
		}

		// Authenticate with the bridge using REAL JWT
		authData := auth.ServiceAccountAuth{
			ClientID:     loginClientID,
			ClientSecret: loginSecret,
			TenantID:     loginTenantID,
		}

		fmt.Printf("Authenticating client: %s\n", authData.ClientID)
		token, err := c.Authenticate(ctx, authData)
		if err != nil {
			return fmt.Errorf("authentication failed: %w", err)
		}

		// Save the REAL JWT token to disk
		if err := auth.SaveToken(token); err != nil {
			return fmt.Errorf("failed to cache token: %w", err)
		}

		fmt.Printf("login: OK (client=%s) — REAL JWT token cached at ~/.helix/token\n", authData.ClientID)
		return nil
	},
}

func init() {
	loginCmd.Flags().StringVar(&loginBaseURL, "base-url", "http://127.0.0.1:3000", "Bridge API base URL")
	loginCmd.Flags().StringVar(&loginClientID, "client-id", "", "Client ID")
	loginCmd.Flags().StringVar(&loginSecret, "client-secret", "", "Client Secret")
	loginCmd.Flags().StringVar(&loginTenantID, "tenant-id", "", "Tenant ID")
	loginCmd.Flags().DurationVar(&loginTimeout, "timeout", 5*time.Second, "Request timeout")

	// Mark required flags
	loginCmd.MarkFlagRequired("client-id")
	loginCmd.MarkFlagRequired("client-secret")
}
