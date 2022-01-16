package cli

import "github.com/spf13/cobra"

func addInitCmds(rootCmd *cobra.Command) {
	s := "Initialise a kbt repository"
	l := "Initialise a kbt repository"

	addIntermediateCmd(rootCmd, "init", s, l)
}
