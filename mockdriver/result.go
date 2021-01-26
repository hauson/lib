package mockdriver

// Result is the result of a query execution.
type Result struct {
	lastInsertID int64
	rowsAffected int64
}

// LastInsertId mock
func (r *Result) LastInsertId() (int64, error) {
	return r.lastInsertID, nil
}

// RowsAffected mock
func (r *Result) RowsAffected() (int64, error) {
	return r.rowsAffected, nil
}
