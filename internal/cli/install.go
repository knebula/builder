package cli

import "github.com/spf13/cobra"

func addInstallCmd(rootCmd *cobra.Command) {
	s := "Deploy a kbt repository"
	l := "Deploy a kbt repository"

	addMageCmd(rootCmd, "install", s, l)
}
