package sqlgenerator

import (
	"testing"
	"fmt"

	sq "github.com/squirrel"
)

func TestGen(t *testing.T) {
	builder := sq.Insert("users").Columns("name", "age").Values("moe", 13).Values("larry", sq.Expr("? + 5", 12))
	sql, err := ToSql(builder)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("sql:", sql)
}
