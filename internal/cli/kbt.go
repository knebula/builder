package cli

import (
	"github.com/spf13/cobra"
)

func NewKbtCmd() (*cobra.Command, error) {
	flags := newFlags()

	s := "The KNebula Build Tool (kbt)"
	l := "The KNebula Build Tool (kbt)"

	kbtCmd := &cobra.Command{
		Use:   "kbt",
		Short: s,
		Long:  l,
	}

	addInitCmd(kbtCmd, flags)
	addInstallCmd(kbtCmd)
	addBuildCmd(kbtCmd)
	addTestCmd(kbtCmd)
	addPackageCmd(kbtCmd)
	addDeployCmd(kbtCmd)

	return kbtCmd, nil
}
