package exchan

import (
	"testing"
	"time"
	"fmt"
)

func TestExChan(t *testing.T) {
	exchan := New()

	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(time.Second)
			exchan.InCh() <- i
		}
	}()

	go func() {
		time.Sleep(7 * time.Second)
		for v := range exchan.OutCh() {
			fmt.Println(v)
		}
	}()

	select {}
}
