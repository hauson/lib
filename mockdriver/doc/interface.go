package mockdriver

/*
import (
	"context"
	"reflect"
)

// DriverContext Driver Context
type DriverContext interface {
	OpenConnector(name string) (Connector, error)
}

// Connector represents a driver in a fixed configuration
type Connector interface {
	Connect(context.Context) (Conn, error)
	Driver() Driver
}

// Pinger is an optional interface that may be implemented by a Conn.
type Pinger interface {
	Ping(ctx context.Context) error
}

// Execer is an optional interface that may be implemented by a Conn.
type Execer interface {
	Exec(query string, args []Value) (Result, error)
}

// ExecerContext is an optional interface that may be implemented by a Conn.
type ExecerContext interface {
	ExecContext(ctx context.Context, query string, args []NamedValue) (Result, error)
}

// Queryer is an optional interface that may be implemented by a Conn.
type Queryer interface {
	Query(query string, args []Value) (Rows, error)
}

// QueryerContext must honor the context timeout and return when the context is canceled.
type QueryerContext interface {
	QueryContext(ctx context.Context, query string, args []NamedValue) (Rows, error)
}

// ConnPrepareContext enhances the Conn interface with context.
type ConnPrepareContext interface {
	PrepareContext(ctx context.Context, query string) (Stmt, error)
}

// ConnBeginTx enhances the Conn interface with context and TxOptions.
type ConnBeginTx interface {
	BeginTx(ctx context.Context, opts TxOptions) (Tx, error)
}

// SessionResetter may be implemented by Conn to allow drivers to reset the
type SessionResetter interface {
	ResetSession(ctx context.Context) error
}

// StmtExecContext enhances the Stmt interface by providing Exec with context.
type StmtExecContext interface {
	ExecContext(ctx context.Context, args []NamedValue) (Result, error)
}

// NamedValueChecker may be optionally implemented by Conn or Stmt. It provides
type NamedValueChecker interface {
	CheckNamedValue(*NamedValue) error
}

// ColumnConverter may be optionally implemented by Stmt if the
type ColumnConverter interface {
	ColumnConverter(idx int) ValueConverter
}

// RowsNextResultSet extends the Rows interface by providing a way to signal
type RowsNextResultSet interface {
	Rows
	HasNextResultSet() bool
	NextResultSet() error
}

// RowsColumnTypeScanType may be implemented by Rows. It should return
type RowsColumnTypeScanType interface {
	Rows
	ColumnTypeScanType(index int) reflect.Type
}

// StmtQueryContext enhances the Stmt interface by providing Query with context.
type StmtQueryContext interface {
	QueryContext(ctx context.Context, args []NamedValue) (Rows, error)
}

// RowsColumnTypeDatabaseTypeName may be implemented by Rows. It should return the
type RowsColumnTypeDatabaseTypeName interface {
	Rows
	ColumnTypeDatabaseTypeName(index int) string
}

// RowsColumnTypeLength may be implemented by Rows. It should return the length
type RowsColumnTypeLength interface {
	Rows
	ColumnTypeLength(index int) (length int64, ok bool)
}

// RowsColumnTypeNullable may be implemented by Rows. The nullable value should
type RowsColumnTypeNullable interface {
	Rows
	ColumnTypeNullable(index int) (nullable, ok bool)
}

// RowsColumnTypePrecisionScale may be implemented by Rows. It should return
type RowsColumnTypePrecisionScale interface {
	Rows
	ColumnTypePrecisionScale(index int) (precision, scale int64, ok bool)
}
*/
