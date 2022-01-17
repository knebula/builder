package cli

import (
	"os"

	"github.com/magefile/mage/mage"
	"github.com/magefile/mage/mg"
	"github.com/spf13/cobra"
)

func addMageCmd(parent *cobra.Command, use, short, long string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		RunE: func(cmd *cobra.Command, args []string) error {
			args = append([]string{use}, args...)
			code := mage.ParseAndRun(os.Stdout, os.Stderr, os.Stdin, args)
			if code != 0 {
				return mg.Fatal(code, args)
			}
			return nil
		},
	}
	if parent != nil {
		parent.AddCommand(cmd)
	}
	return cmd
}
