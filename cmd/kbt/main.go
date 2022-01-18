package main

import (
	"fmt"
	"os"

	"github.com/knebula/builder/internal/cli"
	"github.com/magefile/mage/mg"
)

const version = "v0.0.1"

func main() {
	cmd, err := cli.NewKbtCmd(version)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	if err := cmd.Execute(); err != nil {
		os.Exit(mg.ExitStatus(err))
	}
}
