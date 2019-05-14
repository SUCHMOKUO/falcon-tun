package tun

import (
	"github.com/SUCHMOKUO/falcon-tun/tcpip"
	"log"
)

type PacketHandler = func(*TUN, tcpip.IPv4Packet)
type PacketHandlers = map[tcpip.IPProtocol]PacketHandler

// register packet handlers.
var packetHandlers = PacketHandlers{

	// TCP packet handler.
	tcpip.TCP: func(tun *TUN, ipv4Packet tcpip.IPv4Packet) {
		relayAddr := tun.tcpRelay.addr
		nat := tun.tcpRelay.NAT4
		tcpPacket := tcpip.TCPPacket(ipv4Packet.Payload())
		srcPort := tcpPacket.SourcePort()
		dstPort := tcpPacket.DestinationPort()
		srcIP := ipv4Packet.SourceIP()
		dstIP := ipv4Packet.DestinationIP()

		if srcIP.Equal(relayAddr.IP) && int(srcPort) == relayAddr.Port {
			// it's from relay.
			realSrcPort, realDstIP, realDstPort := nat.GetRecord(dstPort)
			ipv4Packet.SetSourceIP(tun.IP)
			ipv4Packet.SetDestinationIP(realDstIP)
			tcpPacket.SetSourcePort(realSrcPort)
			tcpPacket.SetDestinationPort(realDstPort)
		} else {
			// it's from others, send it to relay.
			natPort := nat.AddRecord(srcPort, dstIP, dstPort)
			ipv4Packet.SetSourceIP(dstIP)
			ipv4Packet.SetDestinationIP(relayAddr.IP)
			tcpPacket.SetSourcePort(natPort)
			tcpPacket.SetDestinationPort(uint16(relayAddr.Port))
		}

		tcpPacket.ResetChecksum(ipv4Packet.PseudoSum())
		ipv4Packet.ResetChecksum()

		_, err := tun.ifce.Write(ipv4Packet)
		if err != nil {
			log.Println("write tun error:", err)
		}
	},

	// ICMP packet handler.
	tcpip.ICMP: func(tun *TUN, ipv4Packet tcpip.IPv4Packet) {
		icmpPacket := tcpip.ICMPPacket(ipv4Packet.Payload())

		if icmpPacket.Type() != tcpip.ICMPRequest || icmpPacket.Code() != 0 {
			return
		}

		log.Printf("ping: %v -> %v\n", ipv4Packet.SourceIP(), ipv4Packet.DestinationIP())

		// forge a reply.
		icmpPacket.SetType(tcpip.ICMPEcho)
		srcIP := ipv4Packet.SourceIP()
		dstIP := ipv4Packet.DestinationIP()
		ipv4Packet.SetSourceIP(dstIP)
		ipv4Packet.SetDestinationIP(srcIP)

		icmpPacket.ResetChecksum()
		ipv4Packet.ResetChecksum()

		_, err := tun.ifce.Write(ipv4Packet)
		if err != nil {
			log.Println("write tun error:", err)
		}
	},
}