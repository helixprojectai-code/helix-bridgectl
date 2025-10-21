package main

import (
	"github.com/helixprojectai-code/helix-bridgectl/internal/cli"
)

var version = "0.1.0"

func main() {
	app := cli.New(version)
	app.Execute()
}
