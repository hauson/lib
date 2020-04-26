package ranker

import (
	"sync"
	"reflect"

	"github.com/emirpasic/gods/trees/redblacktree"
)

type Ranker struct {
	sync.RWMutex
	rbTree *redblacktree.Tree
}

type Value interface {
	Add(Value) Value
}

type Key interface {
	//1: a > other, 0: a == other, -1: a < other
	Cmp(Key) int
}

func New() *Ranker {
	return &Ranker{
		rbTree: redblacktree.NewWith(comparator),
	}
}

func (r *Ranker) Add(key Key, value Value) {
	r.Lock()
	defer r.Unlock()

	r.rbTree.Put(key, value)
}

func (r *Ranker) Delete(key Key) {
	r.Lock()
	defer r.Unlock()

	r.rbTree.Remove(key)
}

func (r *Ranker) Update(key Key, delta Value) {
	r.Lock()
	defer r.Unlock()

	value, ok := r.rbTree.Get(key)
	if !ok {
		r.rbTree.Put(key, delta)
		return
	}

	curScore := value.(Value)
	score := curScore.Add(delta)
	r.rbTree.Put(key, score)
}

func (r *Ranker) Rank(key Key) (rank int, value interface{}) {
	r.RLock()
	defer r.RUnlock()

	v, ok := r.rbTree.Get(key)
	if !ok {
		return 0, nil
	}

	for i, curKey := range r.rbTree.Keys() {
		if reflect.DeepEqual(key, curKey) {
			rank = i
			break
		}
	}

	return rank + 1, v
}

func (r *Ranker) Tops(limit int) (keys, values []interface{}) {
	r.RLock()
	defer r.RUnlock()

	return r.rbTree.Tops(limit)
}

func (r *Ranker) Clear() {
	r.Lock()
	defer r.Unlock()

	r.rbTree.Clear()
}

func comparator(a, b interface{}) int {
	aKey, bKey := a.(Key), b.(Key)
	return aKey.Cmp(bKey)
}
