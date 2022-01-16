package main

import (
	"fmt"
	"os"

	"github.com/knebula/builder/internal/cli"
)

func main() {
	cmd, err := cli.NewKbtCmd()
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
