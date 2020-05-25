package prioqueue

import "container/heap"

// Item of queue
type Item interface {
	Key() string
	Less(Item) bool
}

type prioItem struct {
	Item
	index int
}

// PrioQueue priority queue
type PrioQueue struct {
	items     []*prioItem
	key2Items map[string]*prioItem
}

// New return *PrioQueue
func New(items ...Item) *PrioQueue {
	prioItems := make([]*prioItem, len(items))
	key2Item := make(map[string]*prioItem)
	for i, item := range items {
		if _, ok := key2Item[item.Key()]; ok {
			continue
		}

		prioItems[i] = &prioItem{Item: item, index: i}
		key2Item[item.Key()] = prioItems[i]
	}

	q := &PrioQueue{items: prioItems, key2Items: key2Item}
	heap.Init(q)
	return q
}

// Len length
func (pq *PrioQueue) Len() int {
	return len(pq.items)
}

// Less judge items[i] < items[j]
func (pq *PrioQueue) Less(i, j int) bool {
	return pq.items[i].Less(pq.items[j].Item)
}

// Swap item i and j
func (pq *PrioQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	pq.items[i].index, pq.items[j].index = i, j
}

// Pop last item
func (pq *PrioQueue) Pop() interface{} {
	n := len(pq.items)
	prioItem := pq.items[n-1]
	pq.items = pq.items[0 : n-1]
	delete(pq.key2Items, prioItem.Key())
	return prioItem.Item
}

// Push a item in tail
func (pq *PrioQueue) Push(x interface{}) {
	n := len(pq.items)
	item := x.(Item)
	if _, ok := pq.key2Items[item.Key()]; ok {
		return
	}

	prioItem := &prioItem{Item: item, index: n}
	pq.items = append(pq.items, prioItem)
	pq.key2Items[item.Key()] = prioItem
}

// PopRoot pop root item
func (pq *PrioQueue) PopRoot() (Item, bool) {
	if len(pq.items) == 0 {
		return nil, false
	}

	item := heap.Pop(pq).(Item)
	return item, true
}

// Root return root
func (pq *PrioQueue) Root() (Item, bool) {
	if len(pq.items) == 0 {
		return nil, false
	}

	return pq.items[0].Item, true
}

// Add add item
func (pq *PrioQueue) Add(x Item) {
	heap.Push(pq, x)
}

// Update when priority change
func (pq *PrioQueue) Update(x Item) {
	if v, ok := pq.key2Items[x.Key()]; ok {
		heap.Fix(pq, v.index)
	}
}

// Get item by key
func (pq *PrioQueue) Get(key string) (Item, bool) {
	prioItem, ok := pq.key2Items[key]
	if ok {
		return prioItem.Item, true
	}

	return nil, false
}
