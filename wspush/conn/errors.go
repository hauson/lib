package conn

import "errors"

var ErrClosed = errors.New("conn is closed")
var ErrLogined = errors.New("account is logined")
var ErrNotLogin = errors.New("not logined")
var ErrAcount = errors.New("err account")
