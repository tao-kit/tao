package errorx

import (
	"fmt"
	"strings"

	"manlu.org/tao/tools/taoctl/pkg/env"
)

var errorFormat = `taoctl error: %+v
taoctl env:
%s
%s`

// TaoctlError represents a taoctl error.
type TaoctlError struct {
	message []string
	err     error
}

func (e *TaoctlError) Error() string {
	detail := wrapMessage(e.message...)
	return fmt.Sprintf(errorFormat, e.err, env.Print(), detail)
}

// Wrap wraps an error with taoctl version and message.
func Wrap(err error, message ...string) error {
	e, ok := err.(*TaoctlError)
	if ok {
		return e
	}

	return &TaoctlError{
		message: message,
		err:     err,
	}
}

func wrapMessage(message ...string) string {
	if len(message) == 0 {
		return ""
	}
	return fmt.Sprintf(`message: %s`, strings.Join(message, "\n"))
}
