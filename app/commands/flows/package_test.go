package flows_test

import (
	"bytes"

	"testing"

	"github.com/codegangsta/cli"

	jt "github.com/juju/testing"
	gc "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	gc.TestingT(t)
}

type cmdSuite struct {
	jt.CleanupSuite
	// jt.FakeHomeSuite
	Context *cli.Context
	stdErr  bytes.Buffer
}

var _ = gc.Suite(&cmdSuite{})
