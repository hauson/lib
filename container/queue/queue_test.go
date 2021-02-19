package queue

import (
	"testing"
	"fmt"
)

func TestQueue(t *testing.T) {
	q := NewQueue()
	for i := 8; i < 10; i++ {
		q.EnQueue(i)
	}

	for _, data := range q.All() {
		fmt.Println(data)
	}

	fmt.Println("--------------------------------")

	for _, data := range q.AllInReverse() {
		fmt.Println(data)
	}
}

func TestCircularQueueNotFull(t *testing.T) {
	q := NewCircularQueue(10)
	for i := 1; i < 5; i++ {
		q.EnQueue(i)
	}

	fmt.Println("len:", q.Len(), "cap:", q.Cap())

	items := q.All()
	fmt.Println(items)

	reverseItems := q.AllInReverse()
	fmt.Println(reverseItems)
}

func TestCircularQueueOverFull(t *testing.T) {
	q := NewCircularQueue(3)
	for i := 0; i < 100; i++ {
		q.EnQueue(i)
	}

	fmt.Println("len:", q.Len())

	items := q.All()
	fmt.Println(items)

	reverseItems := q.AllInReverse()
	fmt.Println(reverseItems)
}

func TestReverse(t *testing.T) {
	var s0 []interface{} = nil
	reverse(s0)
	fmt.Println(s0)

	s1 := []interface{}{"a"}
	reverse(s1)
	fmt.Println(s1)

	s2 := []interface{}{"a", "b"}
	reverse(s2)
	fmt.Println(s2)

	s3 := []interface{}{"a", "b", "c"}
	reverse(s3)
	fmt.Println(s3)

	s4 := []interface{}{1, 2, 3, 4, 5}
	reverse(s4)
	fmt.Println(s4)
}
