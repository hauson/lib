package pusher

import (
	"sync"

	"github.com/hauson/lib/wspush/conn"
)

type connects struct {
	sync.RWMutex
	conns map[string]*conn.Connection
}

func newConns() *connects {
	return &connects{
		conns: map[string]*conn.Connection{},
	}
}

func (cs *connects) add(c *conn.Connection) {
	cs.Lock()
	defer cs.Unlock()

	cs.conns[c.IPAddr()] = c
}

func (cs *connects) del(c *conn.Connection) {
	cs.Lock()
	defer cs.Unlock()

	delete(cs.conns, c.IPAddr())
}

func (cs *connects) all() []*conn.Connection {
	cs.RLock()
	defer cs.RUnlock()

	ss := make([]*conn.Connection, 0, len(cs.conns))
	for _, c := range cs.conns {
		ss = append(ss, c)
	}
	return ss
}
