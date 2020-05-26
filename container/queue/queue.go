package queue

import (
	"container/list"
)

//Queue keep to FIFO
type Queue struct {
	*list.List
}

func New() *Queue {
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
	if q.Empty() {
		return nil
	}

	return q.Remove(q.Front())
}

//Empty judge a queue is empty
func (q *Queue) Empty() bool {
	return q.Len() == 0
}
