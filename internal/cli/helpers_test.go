package cli

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

type cmdProxy struct {
	cmd  *cobra.Command
	dir  string
	outb *bytes.Buffer
	errb *bytes.Buffer
}

func newCmdProxy(t *testing.T, f func() *cobra.Command, args []string) *cmdProxy {
	dir := t.TempDir()

	outb := new(bytes.Buffer)
	errb := new(bytes.Buffer)
	args = append([]string{"-d", dir}, args...)

	cmd := f()
	cmd.SetArgs(args)
	cmd.SetOut(outb)
	cmd.SetErr(errb)

	return &cmdProxy{
		cmd:  cmd,
		dir:  dir,
		outb: outb,
		errb: errb,
	}
}

func (c *cmdProxy) execute() error {
	return c.cmd.Execute()
}

func (c *cmdProxy) error(t *testing.T) string {
	out, err := io.ReadAll(c.errb)
	assert.NoError(t, err)
	return string(out)
}

func (c *cmdProxy) cleanup() error {
	return os.RemoveAll(c.dir)
}
