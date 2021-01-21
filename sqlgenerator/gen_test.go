package sqlgenerator

import (
	"fmt"
	"testing"

	"github.com/hauson/lib/mock"
)

func TestInsertSql(t *testing.T) {
	v := mock.MockAdds{Name: "cuihaoxin"}
	sql, err := Sql(Insert, v)
	if err != nil {
		t.Error(err)
	}

	if sql != "INSERT INTO mock_addss (name) VALUES ('cuihaoxin');" {
		t.Error("not equal")
	}
}

func TestDeleteSql(t *testing.T) {
	v := mock.MockAdds{Name: "cuihaoxin"}
	sql, err := Sql(Delete, v)
	if err != nil {
		t.Error(err)
	}

	if sql != "DELETE FROM mock_addss WHERE name = 'cuihaoxin';" {
		t.Error("not equal")
	}
}

func TestUpdateSql(t *testing.T) {
	v := mock.MockAdds{ID: 7, Name: "cuihaoxin"}
	sql, err := Sql(Update, v)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(sql)
}
