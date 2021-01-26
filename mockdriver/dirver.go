package mockdriver

import (
	"strings"
	"fmt"
	"errors"
	"database/sql"
	"database/sql/driver"
)

var ErrNotImplement = errors.New("not implement")

func init() {
	sql.Register("mysql", new(Driver))
}

// Driver mock dirver
type Driver struct{}

// Open mock
func (d *Driver) Open(name string) (driver.Conn, error) {
	ss := strings.Split(name, "/")
	if len(ss) != 2 {
		return nil, fmt.Errorf("%s format err", name)
	}

	subs := strings.Split(ss[1], "?")
	if len(subs) < 0 {
		return nil, fmt.Errorf("%s format err", ss[1])
	}

	return &Conn{dbName: subs[0]}, nil
}

//-------------------------------------------------------------
// Conn is a connection to a database
type Conn struct {
	dbName string
}

func (c *Conn) Prepare(query string) (driver.Stmt, error) {
	return NewStmt(c.dbName, query)
}

func (c *Conn) Close() error {
	return errors.New("Conn Close not implement")
}

func (c *Conn) Begin() (driver.Tx, error) {
	return &Tx{}, nil
}

//----------------------------------------------------------------
// Stmt is a prepared statement
type Stmt struct {
	dbName    string
	tableName string
	numInput  int
	curdMode  CURDMode
	query     string
}

// NewStmt new stmt
func NewStmt(dbName, query string) (*Stmt, error) {
	curdMode, err := parseCurdType(query)
	if err != nil {
		return nil, err
	}

	tableName, err := parseTable(query)
	if err != nil {
		return nil, err
	}

	return &Stmt{
		dbName:    dbName,
		tableName: tableName,
		numInput:  parseNumInput(query),
		curdMode:  curdMode,
		query:     query,
	}, nil
}

func (s *Stmt) NumInput() int {
	return s.numInput
}

func (s *Stmt) Exec(args []driver.Value) (driver.Result, error) {
	db := GetDB(s.dbName)
	table := db.GetTable(s.tableName)
	fileds, err := parseFields(s.query)
	if err != nil {
		return nil, err
	}

	return table.Exec(s.curdMode, fileds, args)
}

func (s *Stmt) Query(args []driver.Value) (driver.Rows, error) {
	//todo:
	db := GetDB(s.dbName)
	table := db.GetTable(s.tableName)
	_ = table

	return &Rows{
		columns: table.fields,
		records:  table.records,
	}, nil
}

func (s *Stmt) Close() error {
	return errors.New("Stmt Close " + ErrNotImplement.Error())
}

//----------------------------------------------------------------------------
// Tx is a transaction.
type Tx struct{}

// Commit mock
func (t *Tx) Commit() error {
	return nil
}

// Rollback mock
func (t *Tx) Rollback() error {
	return nil
}
