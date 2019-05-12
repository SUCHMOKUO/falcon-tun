package util

import "net"

// IPv4ToUint32 convert IPv4 to uint32.
func IPv4ToUint32(ip net.IP) uint32 {
	ip = ip.To4()
	res := uint32(ip[0]) << 24
	res |= uint32(ip[1]) << 16
	res |= uint32(ip[2]) << 8
	res |= uint32(ip[3])
	return res
}

// Uint32ToIPv4 convert uint32 to IPv4.
func Uint32ToIPv4(n uint32) net.IP {
	d := byte(n)
	c := byte(n >> 8)
	b := byte(n >> 16)
	a := byte(n >> 24)
	return net.IPv4(a, b, c, d).To4()
}
