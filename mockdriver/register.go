package mockdriver

import (
	"database/sql"
)

func init() {
	sql.Register("mysql", new(Driver))
}
