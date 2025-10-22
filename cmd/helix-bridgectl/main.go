package main

import (
	"os"

	"github.com/helixprojectai-code/helix-bridgectl/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
