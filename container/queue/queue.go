package queue

import (
	"container/list"
)

//Queue keep to FIFO
type Queue struct {
	*list.List
}

// NewQueue new queue
func NewQueue() *Queue {
	return &Queue{
		List: list.New(),
	}
}

//EnQueue add a data to tail
func (q *Queue) EnQueue(data interface{}) {
	q.PushBack(data)
}

//DeQueue pop a data from head
func (q *Queue) DeQueue() interface{} {
	if q.IsEmpty() {
		return nil
	}

	return q.Remove(q.Front())
}

//IsEmpty judge a queue is empty
func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}

// All get all elements order from first to last, but not delete
func (q *Queue) All() []interface{} {
	items := q.AllInReverse()
	reverse(items)
	return items
}

// AllInReverse get all elements order from last to first, but not delete
func (q *Queue) AllInReverse() []interface{} {
	if q.IsEmpty() {
		return nil
	}

	items := make([]interface{}, 0, q.Len())
	elem := q.Back()
	for elem != nil {
		items = append(items, elem.Value)
		elem = elem.Prev()
	}

	return items
}
