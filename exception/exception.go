package exception

import (
	"runtime"

	"github.com/sirupsen/logrus"
)

//Catch catch exception
func Catch() {
	if err := recover(); err != nil {
		logrus.Errorf("err:%s, stack:%s", err, stack())
	}
}

//Try panic if err is not nil
func Try(fn func() error) {
	if err := fn(); err != nil {
		panic(err)
	}
}

//stack return stack info
func stack() string {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	return string(buf[:n])
}
