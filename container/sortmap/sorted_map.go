package sortmap

import (
	"sort"
	"sync"
)

type Key interface {
	Less(other Key) bool
	Equal(other Key) bool
}

type SortedKeys []Key

func (ks SortedKeys) Len() int           { return len(ks) }
func (ks SortedKeys) Swap(i, j int)      { ks[i], ks[j] = ks[j], ks[i] }
func (ks SortedKeys) Less(i, j int) bool { return ks[i].Less(ks[j]) }

type SortedMap struct {
	sortedKeys SortedKeys
	m          map[Key]interface{}
	sync.RWMutex
}

func NewSortedMap() *SortedMap {
	return &SortedMap{
		m: make(map[Key]interface{}),
	}
}

func (sm *SortedMap) Add(k Key, val interface{}) {
	sm.Lock()
	defer sm.Unlock()

	if _, isExist := sm.m[k]; !isExist {
		sm.sortedKeys = append(sm.sortedKeys, k)
		sort.Sort(sm.sortedKeys)
	}
	sm.m[k] = val
}

func (sm *SortedMap) Del(k Key) {
	sm.Lock()
	defer sm.Unlock()

	for i, item := range sm.sortedKeys {
		if k.Equal(item) {
			sm.sortedKeys = append(sm.sortedKeys[:i], sm.sortedKeys[i+1:]...)
		}
	}
	delete(sm.m, k)
}

func (sm *SortedMap) IsExist(k Key) bool {
	sm.RLock()
	defer sm.RUnlock()

	_, isExist := sm.m[k]
	return isExist
}

func (sm *SortedMap) Get(k Key) (interface{}, bool) {
	sm.RLock()
	defer sm.RUnlock()

	result, isExist := sm.m[k]
	return result, isExist
}

func (sm *SortedMap) Range() []*Item {
	sm.RLock()
	defer sm.RUnlock()

	items := make([]*Item, len(sm.sortedKeys))
	for i, k := range sm.sortedKeys {
		v, _ := sm.m[k]
		items[i] = &Item{k, v}
	}
	return items
}

type Item struct {
	K Key
	V interface{}
}
