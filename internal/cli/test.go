package cli

import "github.com/spf13/cobra"

func addTestCmds(rootCmd *cobra.Command) {
	s := "Test a kbt repository"
	l := "Test a kbt repository"

	addIntermediateCmd(rootCmd, "test", s, l)
}
