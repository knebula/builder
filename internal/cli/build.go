package cli

import "github.com/spf13/cobra"

func addBuildCmd(rootCmd *cobra.Command) {
	s := "Build a kbt repository"
	l := "Build a kbt repository"

	addMageCmd(rootCmd, "build", s, l)
}
