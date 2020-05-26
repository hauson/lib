package wspush

import (
	"sync"
)

// connManager used to manage websocket connection
type connManager struct {
	lock *sync.RWMutex
	//todo: 这个结构
	// address -> account_key -> conn
	connections map[string]map[string]*connection
}

func newConnManager() *connManager {
	return &connManager{
		lock:        new(sync.RWMutex),
		connections: make(map[string]map[string]*connection),
	}
}

func (c *connManager) all() []*connection {
	c.lock.RLock()
	defer c.lock.RUnlock()

	var all []*connection
	for _, conns := range c.connections {
		for _, conn := range conns {
			all = append(all, conn)
		}
	}
	return all
}

func (c *connManager) add(conn *connection) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.delByAccount(conn.loggedAccount)

	conns, ok := c.connections[conn.loggedAccount.address]
	if !ok {
		conns = make(map[string]*connection)
		c.connections[conn.loggedAccount.address] = conns
	}
	conns[conn.loggedAccount.key()] = conn
}

func (c *connManager) del(conn *connection) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.delByAccount(conn.loggedAccount)
}

func (c *connManager) delByAccount(account *account) {
	if conns, ok := c.connections[account.address]; ok {
		delete(conns, account.key())
	}
}

func (c *connManager) getByAccount(account *account) *connection {
	c.lock.RLock()
	defer c.lock.RUnlock()

	if conns, ok := c.connections[account.address]; ok {
		return conns[account.key()]
	}
	return nil
}

func (c *connManager) getByAddress(address string) []*connection {
	c.lock.RLock()
	defer c.lock.RUnlock()

	var conns []*connection
	for _, conn := range c.connections[address] {
		conns = append(conns, conn)
	}
	return conns
}

// todo: 这是一个组件
// todo: 不要通过Server去引用
