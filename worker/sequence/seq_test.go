package sequence

import (
	"testing"
	"fmt"
)

func TestSeq(t *testing.T) {
	sequence := New(100)
	for i := 0; i < 1000; i++ {
		seq := sequence.Next()
		fmt.Println("seq:", seq)
	}
}
