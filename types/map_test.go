package types

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	m1 := map[string]S{
		"aa": {
			Name:   "1",
			Amount: 2,
		},
		"bb": {
			Name:   "2",
			Amount: 3,
		},
	}
	m2 := map[string]S{
		"aa": {
			Name:   "1",
			Amount: 2,
		},
		"cc": {
			Name:   "1",
			Amount: 3,
		},
	}

	out := make(map[string]S)
	MergeMaps(out, m1, m2)
	fmt.Print(out)
}

type S struct {
	Name   string
	Amount int
}

func (s S) Merge(another S) S {
	return S{
		Name:   s.Name,
		Amount: s.Amount + another.Amount,
	}
}

func TestConvMapToSlice(t *testing.T) {
	m := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
		5: 5,
	}
	s := []int{7,100}
	ConvMapToSlice(m, &s)
	t.Log(s)
}
