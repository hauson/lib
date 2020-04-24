package sqlgenerator

import (
	"strings"
	"errors"
	"fmt"
)

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
