package sqlgenerator

import (
	"errors"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"

	"github.com/hauson/lib/char"
	"github.com/hauson/lib/types"
)

type SqlOpr string

const (
	Update SqlOpr = "update"
	Insert SqlOpr = "insert"
	Delete SqlOpr = "delete"
	//Select SqlOpr = "select"
)

func Sql(opr SqlOpr, i interface{}) (string, error) {
	builder, err := Builder(opr, i)
	if err != nil {
		return "", err
	}

	return ToSql(builder)
}

type builder interface {
	ToSql() (string, []interface{}, error)
}

func ToSql(b builder) (string, error) {
	sql, args, err := b.ToSql()
	if err != nil {
		return "", err
	}

	ss := strings.Split(sql, "?")
	if len(ss) != len(args)+1 {
		return "", errors.New("args len err")
	}

	for i, arg := range args {
		ss[i] = ss[i] + fmt.Sprintf("'%v'", arg)
	}

	return strings.Join(ss, "") + ";", nil
}

func Builder(opr SqlOpr, v interface{}) (builder, error) {
	columns, values := columnsAndValues(v)
	switch opr {
	case Insert:
		return sq.Insert(tableName(v)).Columns(columns...).Values(values...), nil
	case Update:
		eq := sq.Eq{}
		clauses := map[string]interface{}{}
		for i, column := range columns {
			if column == "id" {
				eq[column] = values[i]
			} else {
				clauses[column] = values[i]
			}
		}
		return sq.Update(tableName(v)).Where(eq).SetMap(clauses), nil
	case Delete:
		eq := sq.Eq{}
		for i, column := range columns {
			eq[column] = values[i]
		}
		return sq.Delete(tableName(v)).Where(eq), nil
	}

	return nil, nil
}

func tableName(i interface{}) string {
	name := types.ElemType(i).Name()
	return formatUnderline(name) + "s"
}

func columnsAndValues(i interface{}) (columns []string, values []interface{}) {
	fieldNames := types.FieldNames(i)
	for i, value := range types.FieldValues(i) {
		if types.IsNil(value) {
			continue
		}

		if types.IsTypeInitValue(value) {
			continue
		}

		values = append(values, value)
		columns = append(columns, formatUnderline(fieldNames[i]))
	}

	return columns, values
}

func formatUnderline(s string) string {
	sLower := strings.ToLower(s)
	if sLower == "id" {
		return "id"
	}

	if strings.HasSuffix(sLower, "id") {
		sTrim := strings.TrimSuffix(sLower, "id")
		s = s[:len(sTrim)] + "_id"
	}

	var underline []byte
	for j, b := range []byte(s) {
		if char.IsUpper(b) {
			if j != 0 {
				underline = append(underline, '_')
			}
			b = char.ToLower(b)
		}
		underline = append(underline, b)
	}

	return string(underline)
}
