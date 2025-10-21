package commands

import "fmt"

type Command interface {
	Execute(args []string) error
	Help() string
}

type VersionCommand struct {
	Version string
}

func (v *VersionCommand) Execute(args []string) error {
	fmt.Printf("helix-bridgectl version %s\n", v.Version)
	return nil
}

func (v *VersionCommand) Help() string {
	return "Show version information"
}
