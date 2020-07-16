package netdevice

import "net"

type IPInfo struct {
	IP       string
	Mac      net.HardwareAddr
	Hostname string
	Manuf    string
}

type NetDevice struct {
	ipNet        *net.IPNet
	hardwareAddr net.HardwareAddr
	netDevice    string
}

func (n *NetDevice) IPv4() string {
	return n.ipNet.IP.To4().String()
}
