package cli

import "github.com/spf13/cobra"

func newDeployCmd() *cobra.Command {
	s := "Deploy a kbt repository"
	l := "Deploy a kbt repository"

	return newMageCmd("deploy", s, l)
}
