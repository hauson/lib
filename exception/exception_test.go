package exception

import (
	"testing"
	"errors"
)

func TestCatch(t *testing.T) {
	defer Catch()
	panic("this is a panic")
}

func TestTry(t *testing.T) {
	Try(func() error {
		return errors.New("this is error")
	})
}
