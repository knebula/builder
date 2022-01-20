package cli

import "github.com/spf13/cobra"

func newTestCmd() *cobra.Command {
	s := "Test a kbt repository"
	l := "Test a kbt repository"

	return newMageCmd("test", s, l, "")
}
