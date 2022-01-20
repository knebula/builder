package cli

import (
	"errors"
	"os"
	"path/filepath"
	"text/template"

	"github.com/knebula/builder/internal/cli/tmpl"

	"github.com/magefile/mage/mg"
	"github.com/spf13/cobra"
)

const workspaceFile = "kbt.go"
const errorCode = 1

var workspaceTemplate = template.Must(template.New("").Parse(tmpl.Kbt))

func newInitCmd(flags *flags, version string) *cobra.Command {
	s := "Initialise a kbt workspace"
	l := "Initialise a kbt workspace"

	cmd := &cobra.Command{
		Use:   "init",
		Short: s,
		Long:  l,
		RunE: func(cmd *cobra.Command, args []string) error {
			return initWorkspace(flags, version)
		},
	}
	flags.registerInit(cmd)
	return cmd
}

func initWorkspace(flags *flags, version string) error {
	fp := filepath.Join(flags.dir, workspaceFile)
	exists, err := workspaceFileExist(fp)
	if err != nil {
		return workspaceAccessError(err)
	}

	if exists && !flags.force {
		return workspaceInitialisedError()
	}

	f, err := os.Create(fp)
	if err != nil {
		return workspaceInitialiseError(err)
	}
	defer f.Close()

	tmplData := struct {
		Version string
	}{
		Version: version,
	}
	if err := workspaceTemplate.Execute(f, tmplData); err != nil {
		return workspaceInitialiseError(err)
	}

	return nil
}

func workspaceFileExist(fp string) (bool, error) {
	_, err := os.Stat(fp)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func workspaceAccessError(err error) error {
	return mg.Fatalf(errorCode, "error accessing workspace directory or file: %v", err)
}

func workspaceInitialisedError() error {
	return mg.Fatal(errorCode, "workspace already initialised, use --force to reinitialise")
}

func workspaceInitialiseError(err error) error {
	return mg.Fatalf(1, "failed to initialise workspace: %v", err)
}
