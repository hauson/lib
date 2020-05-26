package wspush

import "golang.org/x/sys/unix"

func unname() {
	unix.Accept()
	unix.Open()
}