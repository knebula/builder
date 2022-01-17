package cli

import "github.com/spf13/cobra"

func newBuildCmd() *cobra.Command {
	s := "Build a kbt repository"
	l := "Build a kbt repository"

	return newMageCmd("build", s, l)
}
