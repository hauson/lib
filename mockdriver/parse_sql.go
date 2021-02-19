package mockdriver

import (
	"strings"
	"fmt"
	"errors"
	"regexp"
)

func parseCurdType(query string) (CURDMode, error) {
	query = strings.TrimSpace(query)
	ss := strings.Split(query, " ")
	if len(ss) == 0 {
		return CURDUnkown, fmt.Errorf("%s format err", query)
	}

	switch strings.ToUpper(ss[0]) {
	case CURDInsert.String():
		return CURDInsert, nil
	case CURDSelect.String():
		return CURDSelect, nil
	case CURDUpdate.String():
		return CURDUpdate, nil
	case CURDDelete.String():
		return CURDDelete, nil
	default:
		return CURDUnkown, fmt.Errorf("%s format err", query)
	}
}

var selectTableReg, _ = regexp.Compile("(FROM|from) `([a-zA-Z_]+)`")

func parseTableName(query string) (string, error) {
	curdMode, err := parseCurdType(query)
	if err != nil {
		return "", err
	}

	switch curdMode {
	case CURDInsert:
		ss := strings.Split(query, " ")
		if len(ss) < 3 {
			return "", fmt.Errorf("%s format err", query)
		}

		table := ss[2]
		table = strings.TrimPrefix(table, "`")
		table = strings.TrimSuffix(table, "`")
		return table, nil
	case CURDSelect:
		ss := selectTableReg.FindStringSubmatch(query)
		if len(ss) == 0 {
			return "", fmt.Errorf("%s format err", query)
		}

		return ss[len(ss)-1], nil
	case CURDUpdate:
		return "", errors.New("parseTableName update " + ErrNotImplement.Error())
	case CURDDelete:
		return "", errors.New("parseTableName delete " + ErrNotImplement.Error())
	default:
		return "", errors.New("parseTableName unkown " + ErrNotImplement.Error())
	}
}

func parseNumInput(query string) int {
	return strings.Count(query, "?")
}

var filedsReg, _ = regexp.Compile("`[a-zA-Z_]+`")

func parseFields(query string) ([]string, error) {
	tableName, err := parseTableName(query)
	if err != nil {
		return nil, fmt.Errorf("parseTableName %s:", err)
	}

	var fields []string
	ss := filedsReg.FindAllString(query, -1)
	for _, s := range ss {
		s = strings.TrimPrefix(s, "`")
		s = strings.TrimSuffix(s, "`")
		if s != tableName {
			fields = append(fields, s)
		}
	}

	return fields, nil
}

// contain is ss contain str
func contain(ss []string, tar string) bool {
	for _, item := range ss {
		if item == tar {
			return true
		}
	}
	return false
}
