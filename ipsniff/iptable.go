package ipsniff

import (
	"sync"

	"github.com/hauson/lib/ipsniff/netdevice"
)

type IPTable struct {
	ips map[string]*netdevice.IPInfo
	sync.RWMutex
}

func NewIPTable() *IPTable {
	return &IPTable{
		ips: map[string]*netdevice.IPInfo{},
	}
}

func (t *IPTable) Add(data *netdevice.IPInfo) {
	t.Lock()
	defer t.Unlock()

	if _, ok := t.ips[data.IP]; ok {
		return
	}

	t.ips[data.IP] = data
}

func (t *IPTable) Get(ip string) (*netdevice.IPInfo, bool) {
	t.RLock()
	defer t.RUnlock()

	data, ok := t.ips[ip]
	return data, ok
}

func (t *IPTable) IPList() []string {
	t.RLock()
	defer t.RUnlock()

	ipList := (make([]string, 0, len(t.ips)))
	for ip := range t.ips {
		ipList = append(ipList, ip)
	}
	return ipList
}
