package netdevice

import (
	"net"
	"errors"
)

func LocalNetDevice() (*NetDevice, error) {
	ipNets, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, ipNet := range ipNets {
		if ipData, ok := ipv4(ipNet); ok {
			return ipData, nil
		}
	}

	return nil, errors.New("ca not find net info")
}

func ipv4(it net.Interface) (*NetDevice, bool) {
	addrs, _ := it.Addrs()
	for _, a := range addrs {
		ip, ok := a.(*net.IPNet)
		if ok && !ip.IP.IsLoopback() && ip.IP.To4() != nil {
			return &NetDevice{
				ipNet:        ip,
				hardwareAddr: it.HardwareAddr,
				netDevice:    it.Name,
			}, true
		}
	}
	return nil, false
}


