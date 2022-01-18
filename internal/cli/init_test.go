package cli

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/suite"
)

func TestInitTestSuite(t *testing.T) {
	suite.Run(t, new(InitTestSuite))
}

type InitTestSuite struct {
	suite.Suite

	factory func() *cobra.Command
	flags   *flags
}

func (s *InitTestSuite) SetupTest() {
	s.flags = newFlags()
	s.factory = func() *cobra.Command {
		return newInitCmd(s.flags)
	}
}

func (s *InitTestSuite) TestInitCmd_WhenWorkspaceFileExists_FailsWithNoForceFlag() {
	cp := newCmdProxy(s.T(), s.factory, []string{})
	s.addWorkspaceFile(cp.dir)
	defer cp.cleanup()

	err := cp.execute()

	if s.Error(err) {
		s.Contains(cp.error(s.T()), "workspace already initialised")
	}
}

func (s *InitTestSuite) TestInitCmd_WhenWorkspaceFileExists_SucceedsWithForceFlag() {
	var testCases = []struct {
		flags []string
	}{
		{flags: []string{"-f"}},
		{flags: []string{"--force"}},
	}

	for _, tc := range testCases {
		cp := newCmdProxy(s.T(), s.factory, tc.flags)
		s.addWorkspaceFile(cp.dir)
		defer cp.cleanup()

		err := cp.execute()

		if s.NoError(err) {
			s.Empty(cp.error(s.T()))
		}
	}
}

func (s *InitTestSuite) TestInitCmd_WhenWorkspaceFileDoesNotExists_Succeeds() {
	// If workspace file doesn't exist, the init cmd should
	// succeed with and without the force flag.
	var testCases = []struct {
		flags []string
	}{
		{flags: []string{}},
		{flags: []string{"-f"}},
		{flags: []string{"--force"}},
	}

	for _, tc := range testCases {
		cp := newCmdProxy(s.T(), s.factory, tc.flags)
		defer cp.cleanup()

		err := cp.execute()

		if s.NoError(err) {
			s.Empty(cp.error(s.T()))
		}
	}
}

func (s *InitTestSuite) addWorkspaceFile(dir string) {
	fp := filepath.Join(dir, workspaceFile)
	f, err := os.Create(fp)
	s.NoError(err)

	s.NoError(f.Close())
}
