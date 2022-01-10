package generator

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"manlu.org/tao/tools/taoctl/util/pathx"
)

func TestProtoTmpl(t *testing.T) {
	_ = Clean()
	// exists dir
	err := ProtoTmpl(pathx.MustTempDir())
	assert.Nil(t, err)

	// not exist dir
	dir := filepath.Join(pathx.MustTempDir(), "test")
	err = ProtoTmpl(dir)
	assert.Nil(t, err)
}
