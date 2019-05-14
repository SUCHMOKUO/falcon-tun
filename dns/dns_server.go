package dns

import "net"

// Run start a dns server
// listening at the address provided.
func Run(addr *net.UDPAddr) {
	listener, err := net.ListenUDP("udp4", addr)
	if err != nil {
		panic(err)
	}
}
