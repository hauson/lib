package multigo

import (
	"fmt"
	"testing"
	"time"
)

func TestExecByMultiGo(t *testing.T) {
	const N = 50
	s := make([]Runner, 0, N)
	for i := 0; i < N; i++ {
		s = append(s, &AEntry{
			num: i,
		})
	}

	start1 := time.Now()
	var result []int
	for _, item := range s {
		time.Sleep(50 * time.Millisecond)
		a := item.(*AEntry)
		result = append(result, a.num+a.num)
	}
	fmt.Println("no use concurrent spent time:", time.Since(start1).String())

	start2 := time.Now()
	_, err := ExecByMultiGo(3, s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("use concurrent spent time:", time.Since(start2).String())
}

type AEntry struct {
	num int
}

func (e *AEntry) Exec() (Result, error) {
	time.Sleep(100 * time.Millisecond)
	return &AResult{s: []int{e.num + e.num}}, nil
}

type AResult struct {
	s []int
}

func (r *AResult) Data() interface{} {
	return r.s
}

func (r *AResult) Append(another Result) error {
	nums := another.Data().([]int)
	r.s = append(r.s, nums...)
	return nil
}
