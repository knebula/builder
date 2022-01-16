package cli

import "github.com/spf13/cobra"

func addBuildCmds(rootCmd *cobra.Command) {
	s := "Build a kbt repository"
	l := "Build a kbt repository"

	addIntermediateCmd(rootCmd, "build", s, l)
}
