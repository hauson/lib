package mockdriver

import "errors"

// ErrSkip may be returned by some optional interfaces' methods to
var ErrSkip = errors.New("driver: skip fast-path; continue as if unimplemented")

// ErrBadConn should be returned by a driver to signal to the sql
var ErrBadConn = errors.New("driver: bad connection")

// ErrRemoveArgument may be returned from NamedValueChecker to instruct the
var ErrRemoveArgument = errors.New("driver: remove argument from query")

// ErrNotImplement not implement
var ErrNotImplement = errors.New("not implement")
