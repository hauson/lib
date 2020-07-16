package netdevice

import (
	"net"
	"strconv"
	"bytes"
	"math"
	"strings"

	log "github.com/Sirupsen/logrus"
)

type IP uint32

// 将 IP(uint32) 转换成 可读性IP字符串
func (ip IP) String() string {
	var bf bytes.Buffer
	for i := 1; i <= 4; i++ {
		bf.WriteString(strconv.Itoa(int((ip >> ((4 - uint(i)) * 8)) & 0xff)))
		if i != 4 {
			bf.WriteByte('.')
		}
	}
	return bf.String()
}

// 根据IP和mask换算内网IP范围
func IPRange(ipNet *net.IPNet) []IP {
	ip := ipNet.IP.To4()
	var min, max IP
	for i := 0; i < 4; i++ {
		b := IP(ip[i] & ipNet.Mask[i])
		min += b << ((3 - uint(i)) * 8)
	}
	one, _ := ipNet.Mask.Size()
	max = min | IP(math.Pow(2, float64(32-one))-1)

	// max 是广播地址，忽略
	// i & 0x000000ff  == 0 是尾段为0的IP，根据RFC的规定，忽略
	log.Infof("内网IP范围:%s --- %s", min, max)
	var ips []IP
	for i := min; i < max; i++ {
		if i&0x000000ff != 0 {
			ips = append(ips, i)
		}
	}
	return ips
}

// []byte --> IP
func NewIPByBytes(b []byte) IP {
	return IP(IP(b[0])<<24 + IP(b[1])<<16 + IP(b[2])<<8 + IP(b[3]))
}

// string --> IP
func NewIPByString(s string) IP {
	var b []byte
	for _, i := range strings.Split(s, ".") {
		v, _ := strconv.Atoi(i)
		b = append(b, uint8(v))
	}
	return NewIPByBytes(b)
}

// IPList ，实现了sort的排序接口
type IPList []IP

func (ip IPList) Len() int           { return len(ip) }
func (ip IPList) Swap(i, j int)      { ip[i], ip[j] = ip[j], ip[i] }
func (ip IPList) Less(i, j int) bool { return ip[i] < ip[j] }
