package cli

import "github.com/spf13/cobra"

func addDeployCmd(rootCmd *cobra.Command) {
	s := "Deploy a kbt repository"
	l := "Deploy a kbt repository"

	addMageCmd(rootCmd, "deploy", s, l)
}
