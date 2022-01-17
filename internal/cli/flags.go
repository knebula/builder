package cli

import "github.com/spf13/cobra"

func newFlags() *flags {
	return &flags{}
}

// Keep the short flags unique across all usages in kbt
type flags struct {
	force bool // -f --force
}

func (f *flags) registerInit(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolVarP(&f.force, "force", "f", false, "Force initialisation")
}
