package errors

import (
	"errors"
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	err := errors.New("aa")
	err = Wrap(err, "bb", "cc")
	fmt.Println(err)
}
