package mockdriver

import "database/sql/driver"

// Rows is an iterator over an executed query's results.
type Rows struct{}

func (r *Rows) Columns() []string {
	panic(ErrNotImplement)
}

func (r *Rows) Close() error {
	return ErrNotImplement
}

func (r *Rows) Next(dest []driver.Value) error {
	return ErrNotImplement
}
