package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/helixprojectai-code/helix-sdk-go/pkg/client"
	"github.com/helixprojectai-code/helix-sdk-go/pkg/config"
)

var (
	baseURL string
	timeout time.Duration
)

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "SDK wiring sanity check (calls Client.Ping)",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Build SDK config using the fluent API you exposed
		cfg := config.DefaultConfig().
			WithBaseURL(baseURL).
			WithTimeout(timeout)

		// Construct SDK client
		c := client.New(cfg)

		// Call the SDK's Ping method (no auth needed)
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		if err := c.Ping(ctx); err != nil {
			return fmt.Errorf("ping failed: %w", err)
		}
		fmt.Println("ping: OK")
		return nil
	},
}

func init() {
	pingCmd.Flags().StringVar(&baseURL, "base-url", "http://127.0.0.1:3000", "Bridge API base URL")
	pingCmd.Flags().DurationVar(&timeout, "timeout", 5*time.Second, "Request timeout")
}
