package cli

import "github.com/spf13/cobra"

func addPackageCmd(rootCmd *cobra.Command) {
	s := "Package a kbt repository"
	l := "Package a kbt repository"

	addMageCmd(rootCmd, "package", s, l)
}
