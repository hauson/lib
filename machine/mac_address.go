package machine

import (
	"net"
	"strings"
)

func MacAddress() string {
	inters, err := net.Interfaces()
	if err != nil {
		return ""
	}

	var macs []string
	for _, inter := range inters {
		macs = append(macs, inter.HardwareAddr.String())
	}
	return strings.Join(macs, ",")
}
