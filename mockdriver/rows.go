package mockdriver

import (
	"database/sql/driver"
	"io"
)

// Rows is an iterator over an executed query's results.
type Rows struct {
	columns []string
	records []Record
	i       int
}

// Columns mock
func (r *Rows) Columns() []string {
	return r.columns
}

// Close mock
func (r *Rows) Close() error {
	return nil
}

// Next mock
func (r *Rows) Next(dest []driver.Value) error {
	//return errors.New("Rows Next " + ErrNotImplement.Error())
	if r.i >= len(r.records) {
		return io.EOF
	}

	record := r.records[r.i]
	for i := 0; i < len(dest); i++ {
		dest[i] = record[i]
	}

	r.i++
	return nil
}
