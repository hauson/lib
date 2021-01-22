package mockdriver

// Result is the result of a query execution.
type Result struct{}

func (r *Result) LastInsertId() (int64, error) {
	return 0, ErrNotImplement
}

func (r *Result) RowsAffected() (int64, error) {
	return 0, ErrNotImplement
}
