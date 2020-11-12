package testsuit

import (
	"database/sql"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

const (
	localSetting = "multiStatements=true&charset=utf8&parseTime=true&loc=Local"
	sqlMode      = "SET GLOBAL sql_mode = 'STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION'"
)

// TestDB the struct of the manager of database
type DB struct {
	name string
	db   *gorm.DB
}

// New return one new instance of the db
//account "root:@tcp(127.0.0.1:3306)"
func New(account, preName, sqlPath string) (*DB, error) {
	setting := account + "/?" + localSetting
	name := preName + "_" + uuid.New().String()
	sqlFile, err := file(sqlPath)
	if err != nil {
		return nil, err
	}

	if err := createDB(setting, name, sqlFile); err != nil {
		return nil, err
	}

	db, err := gorm.Open("mysql", setting)
	if err != nil {
		return nil, err
	}

	return &DB{
		name: name,
		db:   db,
	}, nil
}

// Master return *gorm.DB
func (d *DB) Master() *gorm.DB {
	return d.db
}

// ClearTables delete tables from database
func (d *DB) ClearTables(tables []string) error {
	defer func() {
		if err := d.db.Exec("SET FOREIGN_KEY_CHECKS = 1").Error; err != nil {
			panic(err)
		}
	}()

	if err := d.db.Exec("SET FOREIGN_KEY_CHECKS=0").Error; err != nil {
		return err
	}

	for _, table := range tables {
		if err := d.db.Exec("truncate `" + table + "`").Error; err != nil {
			return err
		}
	}

	return nil
}

// Delete drop the database
func (d *DB) Delete() error {
	return d.db.Exec("drop schema `" + d.name + "`").Error
}

func createDB(dns, name, schema string) error {
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.Exec(sqlMode); err != nil {
		return err
	}

	if _, err := db.Exec("create schema `" + name + "`"); err != nil {
		return err
	}

	if _, err = db.Exec("USE `" + name + "`"); err != nil {
		return err
	}

	if _, err := db.Exec(schema); err != nil {
		return err
	}

	return nil
}

func file(path string) (string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
