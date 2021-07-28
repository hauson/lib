package ipsniff

import (
	"time"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/hauson/lib/ipsniff/netdevice"
)

type Sniff struct {
	exitSig  chan int
	ips      *IPTable
	ipCh     chan *netdevice.IPInfo
	localNet *netdevice.NetDevice
}

func New() (*Sniff, error) {
	localNet, err := netdevice.LocalNetDevice()
	if err != nil {
		return nil, err
	}

	return &Sniff{
		exitSig:  make(chan int),
		ips:      NewIPTable(),
		ipCh:     make(chan *netdevice.IPInfo, 100),
		localNet: localNet,
	}, nil
}

func (s *Sniff) Run() {
	if os.Geteuid() != 0 {
		log.Fatal("goscan must run as root.")
	}

	go netdevice.LoopListenARP(s.exitSig, s.ipCh, s.localNet)
	go s.loopPulseSendARP()
	go s.reciveIP()
}

func (s *Sniff) Close() {
	close(s.exitSig)
}

func (s *Sniff) IPList() []string {
	return s.ips.IPList()
}

func (s *Sniff) loopPulseSendARP() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	for {
		go netdevice.SendARP(s.localNet)
		select {
		case <-s.exitSig:
			return
		case <-ticker.C:
		}
	}
}

func (s *Sniff) reciveIP() {
	for {
		select {
		case <-s.exitSig:
			return
		case ipInfo := <-s.ipCh:
			s.ips.Add(ipInfo)
		}
	}
}
