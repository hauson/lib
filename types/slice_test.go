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
