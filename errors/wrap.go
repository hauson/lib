package errors

import (
	"fmt"
	"strings"

	"github.com/hauson/lib/print"
)

// Wrap error
func Wrap(err error, v ...string) error {
	var s string
	if err != nil {
		s = err.Error()
	}

	fileLine := print.FileLineNum()
	info := strings.Join(v, ":")
	return fmt.Errorf("[%s]\n %s:%s", fileLine, s, info)
}
