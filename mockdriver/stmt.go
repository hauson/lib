package mockdriver

import (
	"database/sql/driver"
)

// Stmt is a prepared statement
type Stmt struct{}

func (s *Stmt) NumInput() int {
	panic(ErrNotImplement)
}

func (s *Stmt) Exec(args []driver.Value) (driver.Result, error) {
	return nil, ErrNotImplement
}

func (s *Stmt) Query(args []driver.Value) (driver.Rows, error) {
	return nil, ErrNotImplement
}
