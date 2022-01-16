package cli

import "github.com/spf13/cobra"

func addDeployCmds(rootCmd *cobra.Command) {
	s := "Deploy a kbt repository"
	l := "Deploy a kbt repository"

	addIntermediateCmd(rootCmd, "deploy", s, l)
}
