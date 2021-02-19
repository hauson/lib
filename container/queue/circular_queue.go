package queue

import "container/ring"

// CircularQueue circular queue
type CircularQueue struct {
	*ring.Ring
	len int
}

// NewCircularQueue make a c
func NewCircularQueue(n int) *CircularQueue {
	return &CircularQueue{
		Ring: ring.New(n),
	}
}

//EnQueue add a data to tail
func (q *CircularQueue) EnQueue(data interface{}) {
	q.Ring.Value = data
	q.Ring = q.Ring.Next()
	if q.len < q.Cap() {
		q.len++
	}
}

//IsEmpty judge a queue is empty
func (q *CircularQueue) IsEmpty() bool {
	return q.Len() == 0
}

func (q *CircularQueue) Len() int {
	return q.len
}

func (q *CircularQueue) Cap() int {
	return q.Ring.Len()
}

// All get all elements order from first to last, but not delete
func (q *CircularQueue) All() []interface{} {
	items := q.AllInReverse()
	reverse(items)
	return items
}

// AllInReverse get all elements order from last to first, but not delete
func (q *CircularQueue) AllInReverse() []interface{} {
	if q.IsEmpty() {
		return nil
	}

	items := make([]interface{}, 0, q.Len())
	elem := q.Ring.Prev()
	for i := 0; i < q.Len(); i++ {
		items = append(items, elem.Value)
		elem = elem.Prev()
	}

	return items
}

func reverse(s []interface{}) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
