package cli

import "github.com/spf13/cobra"

func newPackageCmd() *cobra.Command {
	s := "Package a kbt repository"
	l := "Package a kbt repository"

	return newMageCmd("package", s, l)
}
