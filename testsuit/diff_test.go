package testsuit

import (
	"encoding/json"
	"testing"
	"time"
)

type Obj struct {
	Name   string
	Number int
	Time   time.Time
	E      Ele
}

type Ele struct {
	Name string
	Num  int
}

func TestDiffObj(t *testing.T) {
	o1 := &Obj{
		Name:   "chx",
		Number: 7,
		Time:   time.Now(),
	}
	o2 := &Obj{
		Name:   "chx",
		Number: 7,
		Time:   time.Now(),
		E: Ele{
			Name: "name2",
			Num:  2,
		},
	}

	diff := DiffObj(o1, o2, "Time", "E")

	t.Log("diff:", diff)
}

func TestDiffSlice(t *testing.T) {
	s1 := []*Ele{
		{Name: "name1", Num: 1},
		{Name: "name2", Num: 3},
	}
	s2 := []*Ele{
		{Name: "name1", Num: 1},
		{Name: "name2", Num: 2},
	}

	diff := DiffSlice(s1, s2)
	t.Log("diff:", diff)
}

func TestDiffMap(t *testing.T) {
	m1 := map[string]interface{}{
		"a": &Ele{Name: "a_name", Num: 1},
		"c": &Ele{Name: "c_name", Num: 2},
	}
	m2 := map[string]interface{}{
		"d": &Ele{Name: "a_name", Num: 1},
		"b": &Ele{Name: "b_name", Num: 2},
	}

	diff := DiffMap(m1, m2)
	t.Log("diff:", diff)
}

func TestClearFields(t *testing.T) {
	a := &Obj{
		Name:   "name1",
		Number: 7,
		E: Ele{
			Name: "ptr_name",
			Num:  71,
		},
	}

	clearFields(a, "E")
	bytes, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("cur:", string(bytes))
}
