package nat

import "net"

func ipv4ToUint32(ip net.IP) uint32 {
	ip = ip.To4()
	res := uint32(ip[0]) << 24
	res |= uint32(ip[1]) << 16
	res |= uint32(ip[2]) << 8
	res |= uint32(ip[3])
	return res
}

func uint32ToIPv4(n uint32) net.IP {
	d := byte(n)
	c := byte(n >> 8)
	b := byte(n >> 16)
	a := byte(n >> 24)
	return net.IPv4(a, b, c, d).To4()
}

func recordToUint64(srcPort uint16, dstIP net.IP, dstPort uint16) uint64 {
	dstIP = dstIP.To4()
	res := uint64(srcPort) << 48
	res |= uint64(ipv4ToUint32(dstIP)) << 16
	res |= uint64(dstPort)
	return res
}

func uint64ToRecord(n uint64) (srcPort uint16, dstIP net.IP, dstPort uint16) {
	dstPort = uint16(n)
	dstIP = uint32ToIPv4(uint32(n >> 16))
	srcPort = uint16(n >> 48)
	return
}