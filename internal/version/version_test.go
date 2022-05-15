package version_test

import (
	"bufio"
	"github.com/ilijamt/blacklist_checker/internal/version"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestPrintVersion(t *testing.T) {
	wr := bufio.NewWriter(ioutil.Discard)
	version.PrintVersion(wr)
	assert.Equal(t, 71, wr.Buffered())
	assert.NoError(t, wr.Flush())
}
