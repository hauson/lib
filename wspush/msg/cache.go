package msg

import "sync"

type Cache struct {
	topics map[string]Raw
	sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		topics: make(map[string]Raw),
	}
}

func (t *Cache) Get(topic string) (Raw, bool) {
	t.RLock()
	defer t.RUnlock()

	msg, ok := t.topics[topic]
	return msg, ok
}

func (t *Cache) Set(raw Raw) {
	t.Lock()
	defer t.Unlock()

	t.topics[raw.Topic()] = raw
}
