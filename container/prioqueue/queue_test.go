package prioqueue

import (
	"testing"
	"fmt"
)

type MyItem struct {
	v      interface{}
	expire int
}

func (m *MyItem) Key() string {
	return fmt.Sprintf("%v", m.v)
}

func (m *MyItem) Less(x Item) bool {
	other := x.(*MyItem)
	return other.expire < m.expire
}

func TestName(t *testing.T) {
	ss := []*MyItem{
		{"二毛", 5,},
		{"张三", 3,},
		{"狗蛋", 9,},
	}
	st := make([]Item, len(ss))
	for i, item := range ss {
		st[i] = item
	}

	pq := New(st...)
	item := &MyItem{
		v:      "李四",
		expire: 4,
	}
	pq.Add(item)
	item.expire = 1
	pq.Update(item)
	for pq.Len() > 0 {
		root, _ := pq.Root()
		r := root.(*MyItem)
		fmt.Printf("v:%v,expire:%v\n", r.v, r.expire)

		v, _ := pq.PopRoot()
		myItem := v.(*MyItem)
		fmt.Printf("v:%v,expire:%v\n", myItem.v, myItem.expire)
	}
}
