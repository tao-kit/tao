package errorx

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrap(t *testing.T) {
	err := errors.New("foo")
	err = Wrap(err)
	_, ok := err.(*TaoctlError)
	assert.True(t, ok)
}
