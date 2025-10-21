package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bridgectl",
	Short: "Aletheia Bridge CLI",
	Long:  "Operations CLI for Aletheia Bridge.",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// placeholder default command
	rootCmd.AddCommand(healthCmd)
	// global flags can be added here as needed
	fmt.Sprintf("") // keep fmt imported if unused
}
