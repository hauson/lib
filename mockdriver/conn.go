package mockdriver

import (
	"database/sql/driver"
)

// Conn is a connection to a database
type Conn struct{}

func (c *Conn) Prepare(query string) (driver.Stmt, error) {
	return nil, ErrNotImplement
}

func (c *Conn) Close() error {
	return ErrNotImplement
}

func (c *Conn) Begin() (driver.Tx, error) {
	return nil, ErrNotImplement
}
