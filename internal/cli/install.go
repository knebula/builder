package cli

import "github.com/spf13/cobra"

func newInstallCmd() *cobra.Command {
	s := "Deploy a kbt repository"
	l := "Deploy a kbt repository"

	return newMageCmd("install", s, l, "")
}
