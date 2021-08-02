package types

import (
	"fmt"
	"testing"
)

func TestSlicePluck(t *testing.T) {
	{
		// case0 s0 is nil
		var s0 []*AA
		var s1 []string
		err := SlicePluck(&s0, "Name", &s1)
		if err != nil {
			panic(err)
		}
	}

	{
		// case0 s0 is ptr slice
		s0 := []*AA{
			{"cui", 1, Sub{10, "hebei"}},
			{"hao", 2, Sub{9, "tainjin"}},
			{"xin", 3, Sub{7, "nanjing"}},
		}

		var s1 []string
		if err := SlicePluck(s0, "Name", &s1); err != nil {
			panic(err)
		}

		fmt.Println(s1)
	}

	{
		// case0 s0 is elem slice
		s0 := []AA{
			{"cui", 1, Sub{10, "hebei"}},
			{"hao", 2, Sub{9, "tainjin"}},
			{"xin", 3, Sub{7, "nanjing"}},
		}

		var s1 []int
		if err := SlicePluck(s0, "Num", &s1); err != nil {
			panic(err)
		}

		fmt.Println(s1)
	}

	{
		// case0 s1 is field struct
		s0 := []AA{
			{"cui", 1, Sub{10, "hebei"}},
			{"hao", 2, Sub{9, "tainjin"}},
			{"xin", 3, Sub{7, "nanjing"}},
		}

		var s1 []Sub
		if err := SlicePluck(s0, "Sub", &s1); err != nil {
			panic(err)
		}

		fmt.Println(s1)
	}
}

type AA struct {
	Name string
	Num  int
	Sub  Sub
}

type Sub struct {
	Sex  int
	Addr string
}
type Obj struct {
	B int
	A string
	C Sub
}

func TestStrings(t *testing.T) {
	{
		t.Log("case : elem slice")
		s0 := []Obj{
			{A: "cui", B: 1, C: Sub{10, "hebei"}},
			{A: "hao", B: 2, C: Sub{9, "tainjin"}},
			{A: "xin", B: 3, C: Sub{7, "nanjing"}},
		}

		lines, err := Strings(s0)
		if err != nil {
			t.Fatal(err)
		}

		for i, line := range lines {
			fmt.Println(i+1, line)
		}
	}

	{
		t.Log("elem ptr slice")
		s0 := []*Obj{
			{A: "cui", B: 1, C: Sub{10, "hebei"}},
			{A: "hao", B: 2, C: Sub{9, "tainjin"}},
			{A: "xin", B: 3, C: Sub{7, "nanjing"}},
		}

		lines, err := Strings(s0)
		if err != nil {
			t.Fatal(err)
		}

		for i, line := range lines {
			fmt.Println(i+1, line)
		}
	}
}

type Entry struct {
	key   string
	value int
}
type Key struct {
	K string
}

type Value struct {
	V int
}

func (e *Entry) Key() *Key {
	return &Key{
		K: e.key,
	}
}

func (e *Entry) Value() *Value {
	return &Value{
		V: e.value,
	}
}

func TestConvSliceToMap(t *testing.T) {
	s := []*Entry{
		{"a", 1},
		{"b", 2},
		{"c", 3},
	}

	m := make(map[*Key]*Value)
	ConvSliceToMap(s, m)
	for k, v := range m {
		fmt.Println(*k, *v)
	}
}
