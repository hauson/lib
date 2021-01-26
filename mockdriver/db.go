package mockdriver

import (
	"sync"
	"errors"
	"database/sql/driver"
	"fmt"
)

var lock sync.Mutex
var dbs = map[string]*DB{}

// GetDB find db
func GetDB(name string) *DB {
	lock.Lock()
	defer lock.Unlock()

	db, ok := dbs[name]
	if !ok {
		db = NewDB(name)
		dbs[name] = db
	}
	return db
}

//-------------------------------------------------------------------
// DB mock
type DB struct {
	name   string
	mu     sync.Mutex
	tables map[string]*Table
}

// NewDB mock
func NewDB(name string) *DB {
	return &DB{
		name:   name,
		tables: make(map[string]*Table),
	}
}

// GetTable  mock
func (db *DB) GetTable(name string) *Table {
	db.mu.Lock()
	defer db.mu.Unlock()

	table, ok := db.tables[name]
	if !ok {
		table = NewTable(name)
		db.tables[name] = table
	}

	return table
}

//------------------------------------------------------------------
// Record mock
type Record []driver.Value

// Table mock
type Table struct {
	name    string
	mu      sync.RWMutex
	fields  []string
	records []Record
}

// NewTable new table
func NewTable(name string) *Table {
	return &Table{
		name: name,
	}
}

// Exec mock
func (t *Table) Exec(mode CURDMode, fields []string, args []driver.Value) (driver.Result, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if len(t.fields) == 0 {
		t.fields = fields
	}

	switch mode {
	case CURDInsert:
		if len(args) != len(t.fields) {
			return nil, fmt.Errorf("args %d not equal fields %d ", len(args), len(t.fields))
		}

		t.records = append(t.records, args)

		return &Result{
			lastInsertID: int64(len(t.records)),
			rowsAffected: 1,
		}, nil
	}

	return nil, errors.New("Table Exec not implement")
}
