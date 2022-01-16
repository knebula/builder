package cli

import "github.com/spf13/cobra"

func addPackageCmds(rootCmd *cobra.Command) {
	s := "Package a kbt repository"
	l := "Package a kbt repository"

	addIntermediateCmd(rootCmd, "package", s, l)
}
