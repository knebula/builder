package cli

import "github.com/spf13/cobra"

func NewKbtCmd() (*cobra.Command, error) {
	s := "The KNebula Build Tool (kbt)"
	l := "The KNebula Build Tool (kbt)"

	kbtCmd := addIntermediateCmd(nil, "kbt", s, l)
	addInitCmds(kbtCmd)
	addBuildCmds(kbtCmd)
	addTestCmds(kbtCmd)
	addPackageCmds(kbtCmd)
	addDeployCmds(kbtCmd)

	return kbtCmd, nil
}

func addIntermediateCmd(parent *cobra.Command, use, short, long string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
	}
	if parent != nil {
		parent.AddCommand(cmd)
	}
	return cmd
}
