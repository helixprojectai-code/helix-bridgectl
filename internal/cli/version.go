package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	version = "v0.1.0"
	commit  = "dev"
	date    = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("helix-bridgectl %s (commit %s, built %s)\n", version, commit, date)
	},
}
