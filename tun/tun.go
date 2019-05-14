package tun

import (
	"github.com/SUCHMOKUO/falcon-tun/dns"
	"github.com/SUCHMOKUO/falcon-tun/relay"
	"github.com/SUCHMOKUO/falcon-tun/tcpip"
	"github.com/songgao/water"
	"net"
	"sync"
)

const mtu = 1500

var bufPool = &sync.Pool{
	New: func() interface{} {
		return make([]byte, mtu)
	},
}

// TUN represent a tun service.
type TUN struct {
	ifce *water.Interface
	tcpRelay *relay.TCPRelay
	Name string
	IP net.IP
}

// Config provide infos to create a tun device.
type Config struct {
	// tun device's ip address.
	IP net.IP
	HandleTCPConn relay.TCPConnHandler
}

// New create the tun service and return the instance.
func New(cfg *Config) *TUN {
	tun := new(TUN)
	// set interface.
	ifce, err := water.New(water.Config{
		DeviceType: water.TUN,
	})
	if err != nil {
		panic(err)
	}
	// set tcp relay.
	tun.tcpRelay = relay.NewTCPRelay(&net.TCPAddr{
		IP: tun.IP,
		Port: 82,
	})
	tun.ifce = ifce
	tun.Name = ifce.Name()
	tun.IP = cfg.IP
	return tun
}

// Run start the tun service.
func (tun *TUN) Run() {
	go dns.Run(&net.UDPAddr{
		IP: tun.IP,
		Port: 53,
	})
	go tun.tcpRelay.Run()

	for {
		buf := bufPool.Get().([]byte)
		n, err := tun.ifce.Read(buf)
		if err != nil {
			bufPool.Put(buf)
			continue
		}
		tun.dispatch(buf[:n])
		bufPool.Put(buf)
	}
}

// dispatch the packet to it's handler.
func (tun *TUN) dispatch(packet []byte) {
	if !tcpip.IsIPv4(packet) {
		return
	}
	ipv4Packet := tcpip.IPv4Packet(packet)
	protocol := ipv4Packet.Protocol()
	handlePacket, exists := packetHandlers[protocol]
	if !exists {
		return
	}
	handlePacket(tun, ipv4Packet)
}