package netdevice

import (
	"time"
	"net"
	log "github.com/sirupsen/logrus"
	"github.com/timest/gomanuf"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

const (
	arpReq  = 0x0001
	arpResp = 0x0002
)

var defaultOpt gopacket.SerializeOptions

//context 信号参数，output 参数
func LoopListenARP(exitSig chan int, ipInfos chan<- *IPInfo, localIP *NetDevice) {
	handle, err := pcap.OpenLive(localIP.netDevice, 1024, false, 10*time.Second)
	if err != nil {
		log.Fatal("pcap打开失败:", err)
	}
	defer handle.Close()

	handle.SetBPFFilter("arp")
	ps := gopacket.NewPacketSource(handle, handle.LinkType())
	for {
		if listenARP(exitSig, ps, ipInfos) {
			return
		}
	}
}

func listenARP(exitSig chan int, ps *gopacket.PacketSource, ipInfos chan<- *IPInfo) (isDone bool) {
	select {
	case <-exitSig:
		return true
	case p := <-ps.Packets():
		if arp := p.Layer(layers.LayerTypeARP).(*layers.ARP); arp.Operation == 2 {
			mac := net.HardwareAddr(arp.SourceHwAddress)
			ipInfos <- &IPInfo{
				IP:       NewIPByBytes(arp.SourceProtAddress).String(),
				Mac:      mac,
				Hostname: "",
				Manuf:    manuf.Search(mac.String()),
			}
		}

		return false
	}
}

func SendARP(localIP *NetDevice) {
	for _, ip := range IPRange(localIP.ipNet) {
		go sendArpPackage(localIP, ip)
	}
}

// 发送arp包, buildPackage 和 send 两个函数
func sendArpPackage(src *NetDevice, dst IP) {
	handle, err := pcap.OpenLive(src.netDevice, 2048, false, 3*time.Second)
	if err != nil {
		log.Fatal("pcap打开失败:", err)
	}
	defer handle.Close()

	packet := buildArpPackage(src, dst)
	if err := handle.WritePacketData(packet.Bytes()); err != nil {
		log.Fatal("发送arp数据包失败..")
	}
}

func buildArpPackage(src *NetDevice, dst IP) gopacket.SerializeBuffer {
	srcIp := net.ParseIP(src.ipNet.IP.String()).To4()
	dstIp := net.ParseIP(dst.String()).To4()
	if srcIp == nil || dstIp == nil {
		log.Fatal("ip 解析出问题")
	}

	ether := ethernetProto(src.hardwareAddr)
	arp := arpProto(src.hardwareAddr, srcIp, dstIp)
	packet := gopacket.NewSerializeBuffer()
	gopacket.SerializeLayers(packet, defaultOpt, ether, arp)

	return packet
}

func arpProto(srcHardwareAddr net.HardwareAddr, srcIP, dstIP net.IP) *layers.ARP {
	return &layers.ARP{
		AddrType:          layers.LinkTypeEthernet,
		Protocol:          layers.EthernetTypeIPv4,
		HwAddressSize:     uint8(6),
		ProtAddressSize:   uint8(4),
		Operation:         arpReq,
		SourceHwAddress:   srcHardwareAddr,
		SourceProtAddress: srcIP,
		DstHwAddress:      net.HardwareAddr{0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		DstProtAddress:    dstIP,
	}
}

func ethernetProto(srcHardwareAddr net.HardwareAddr) *layers.Ethernet {
	return &layers.Ethernet{
		SrcMAC:       srcHardwareAddr,
		DstMAC:       net.HardwareAddr{0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		EthernetType: layers.EthernetTypeARP,
	}
}
