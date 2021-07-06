package batcher

import (
	"fmt"
	"testing"
)

func TestAA(t *testing.T) {
	var s []int
	for i := 0; i < 100; i++ {
		s = append(s, i)
	}

	const batch = 7
	for i := 0; i < len(s); i += batch {
		var page []int
		FillPage(&page, s, i, batch)
		fmt.Println("got:", page)
	}

}
