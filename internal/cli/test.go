package cli

import "github.com/spf13/cobra"

func addTestCmd(rootCmd *cobra.Command) {
	s := "Test a kbt repository"
	l := "Test a kbt repository"

	addMageCmd(rootCmd, "test", s, l)
}
