package tcpip

import (
	"encoding/binary"
	"net"
)

type IPProtocol byte

const (
	ICMP IPProtocol = 0x01
	TCP             = 0x06
	UDP             = 0x11
)

func getIPVersion(packet []byte) byte {
	if len(packet) < 20 {
		return 0
	}
	return packet[0] >> 4
}

func IsIPv4(packet []byte) bool {
	return 4 == getIPVersion(packet)
}

func IsIPv6(packet []byte) bool {
	return 6 == getIPVersion(packet)
}

type IPv4Packet []byte

func (p IPv4Packet) TotalLen() uint16 {
	return binary.BigEndian.Uint16(p[2:])
}

func (p IPv4Packet) HeaderLen() uint16 {
	return uint16(p[0]&0xf) * 4
}

func (p IPv4Packet) DataLen() uint16 {
	return p.TotalLen() - p.HeaderLen()
}

func (p IPv4Packet) Payload() []byte {
	return p[p.HeaderLen():p.TotalLen()]
}

func (p IPv4Packet) Protocol() IPProtocol {
	return IPProtocol(p[9])
}

func (p IPv4Packet) SourceIP() net.IP {
	return net.IPv4(p[12], p[13], p[14], p[15]).To4()
}

func (p IPv4Packet) SetSourceIP(ip net.IP) {
	ip = ip.To4()
	if ip != nil {
		copy(p[12:16], ip)
	}
}

func (p IPv4Packet) DestinationIP() net.IP {
	return net.IPv4(p[16], p[17], p[18], p[19]).To4()
}

func (p IPv4Packet) SetDestinationIP(ip net.IP) {
	ip = ip.To4()
	if ip != nil {
		copy(p[16:20], ip)
	}
}

func (p IPv4Packet) Checksum() uint16 {
	return binary.BigEndian.Uint16(p[10:])
}

func (p IPv4Packet) SetChecksum(sum [2]byte) {
	p[10] = sum[0]
	p[11] = sum[1]
}

func (p IPv4Packet) ResetChecksum() {
	p.SetChecksum(zeroChecksum)
	p.SetChecksum(Checksum(0, p[:p.HeaderLen()]))
}

// for tcp checksum
func (p IPv4Packet) PseudoSum() uint32 {
	sum := Sum(p[12:20])
	sum += uint32(p.Protocol())
	sum += uint32(p.DataLen())
	return sum
}
