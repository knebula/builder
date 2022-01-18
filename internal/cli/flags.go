package cli

import "github.com/spf13/cobra"

func newFlags() *flags {
	return &flags{}
}

// Keep the short flags unique across all usages in kbt
type flags struct {
	force bool   // -f --force
	dir   string // -d --dir
}

func (f *flags) registerInit(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolVarP(&f.force, "force", "f", false, "force recreation of kbt init files")
	cmd.PersistentFlags().StringVarP(&f.dir, "dir", "d", "", "directory to read workspace files from (default \".\")")
}
