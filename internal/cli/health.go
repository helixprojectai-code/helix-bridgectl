package cli

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

type healthResult struct {
	Component string        `json:"component"`
	Status    string        `json:"status"`
	LatencyMS time.Duration `json:"latency_ms"`
}

var output string

var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "Check basic CLI health",
	RunE: func(cmd *cobra.Command, args []string) error {
		// stubbed success; extend to real checks later
		res := []healthResult{{
			Component: "cli",
			Status:    "healthy",
			LatencyMS: 1,
		}}
		switch output {
		case "json":
			b, _ := json.MarshalIndent(res, "", "  ")
			fmt.Println(string(b))
		default:
			fmt.Println("cli: healthy")
		}
		return nil
	},
}

func init() {
	healthCmd.Flags().StringVarP(&output, "output", "o", "table", "output format (table|json)")
}
