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
	if err := MergeMaps(out, m1, m2); err != nil {
		t.Fatal(err)
	}

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
