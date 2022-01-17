package cli

import (
	"errors"
	"os"
	"text/template"

	"github.com/magefile/mage/mg"
	"github.com/spf13/cobra"
)

const (
	workspaceFile = "kbt.go"
)

var (
	workspaceTemplate = template.Must(template.New("").Parse(kbtTmpl))
)

func addInitCmd(rootCmd *cobra.Command, flags *flags) {
	s := "Initialise a kbt repository"
	l := "Initialise a kbt repository"

	cmd := &cobra.Command{
		Use:   "init",
		Short: s,
		Long:  l,
		RunE: func(cmd *cobra.Command, args []string) error {
			return generateWorkspaceFile(flags)
		},
	}
	flags.registerInit(cmd)
	rootCmd.AddCommand(cmd)
}

func generateWorkspaceFile(flags *flags) error {
	exists, err := doesWorkspaceFileExist()
	if err != nil {
		return mg.Fatalf(1, "failed to determine if the kbt workspace template exists: %v", err)
	}

	if exists && !flags.force {
		return mg.Fatal(1, "kbt workspace already initialised, use --force to reinitialise")
	}

	f, err := os.Create(workspaceFile)
	if err != nil {
		return mg.Fatalf(1, "failed to create kbt workspace template: %v", err)
	}
	defer f.Close()

	if err := workspaceTemplate.Execute(f, nil); err != nil {
		return mg.Fatalf(1, "failed to execute kbt workspace template: %v", err)
	}

	return nil
}

func doesWorkspaceFileExist() (bool, error) {
	_, err := os.Stat(workspaceFile)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}
