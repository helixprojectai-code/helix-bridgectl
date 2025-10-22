package cli

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "helix-bridgectl",
	Short: "Helix Bridge CLI",
	Long:  "Command-line interface for Helix Bridge operations.",
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(pingCmd)
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(whoamiCmd) // ⬅️ add this line
}

func Execute() error {
	return rootCmd.Execute()
}
