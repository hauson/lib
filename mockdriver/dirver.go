package mockdriver

import (
	"database/sql/driver"
)

// Driver is the interface that must be implemented by a database
type Driver struct{}

func (d *Driver) Open(name string) (driver.Conn, error) {
	return nil, ErrNotImplement
}
