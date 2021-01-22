package mockdriver

// Tx is a transaction.
type Tx struct{}

func (t *Tx) Commit() error {
	return ErrNotImplement
}

func (t *Tx) Rollback() error {
	return ErrNotImplement
}
