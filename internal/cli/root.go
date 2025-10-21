package cli

import (
	"fmt"
	"os"
)

type CLI struct {
	Version string
}

func New(version string) *CLI {
	return &CLI{
		Version: version,
	}
}

func (c *CLI) Execute() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--version", "-v":
			fmt.Printf("helix-bridgectl version %s\n", c.Version)
			return
		case "help", "--help", "-h":
			c.showHelp()
			return
		}
	}
	
	c.showHelp()
}

func (c *CLI) showHelp() {
	fmt.Printf("Helix Bridge CLI v%s\n", c.Version)
	fmt.Println("\nUsage:")
	fmt.Println("  helix-bridgectl [command]")
	fmt.Println("\nCommands:")
	fmt.Println("  version    Show version information")
	fmt.Println("  help       Show this help message")
	fmt.Println("\nFlags:")
	fmt.Println("  -h, --help     Show help")
	fmt.Println("  -v, --version  Show version")
}
