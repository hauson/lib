package errors

import (
	"fmt"
	"strings"

	"github.com/lib/print"
)

// Wrap error
func Wrap(err error, v ...string) error {
	var s string
	if err != nil {
		s = err.Error()
	}

	file, line := print.FileLine()
	info := strings.Join(v, ":")
	return fmt.Errorf("[%s:%s]\n %s:%s", file, line, s, info)
}
