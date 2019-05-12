package relay

import (
	"github.com/SUCHMOKUO/falcon-tun/dns"
	"github.com/SUCHMOKUO/falcon-tun/nat"
	"io"
	"log"
	"net"
	"strconv"
)

// ConnHandler is the type alias of connection handler function.
type ConnHandler = func(host, port string, conn io.ReadWriteCloser)

// TCPRelay represent a relay for tcp connection.
type TCPRelay struct {
	NAT4 *nat.NAT4
	Handler ConnHandler
	Address string
}

// Run create a new tcp relay.
func (tr *TCPRelay) Run() {
	l, err := net.Listen("tcp4", tr.Address)
	if err != nil {
		log.Fatalln("TCP relay error:", err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}
		host, port := tr.getTargetInfo(conn)
		go tr.Handler(host, port, conn)
	}
}

func (tr *TCPRelay) getTargetInfo(conn net.Conn) (domain, port string) {
	addr := conn.RemoteAddr().String()
	ipStr, natPortStr, err := net.SplitHostPort(addr)
	if err != nil {
		log.Fatalln("Get remote address error:", err)
	}
	natPort, err := strconv.Atoi(natPortStr)
	if err != nil {
		log.Fatalln("Get remote address error:", err)
	}
	// get target port.
	_, _, dstPort := tr.NAT4.GetRecord(uint16(natPort))
	port = strconv.Itoa(int(dstPort))
	// get target domain.
	domain, err = dns.GetDomain(ipStr)
	if err != nil {
		log.Fatalln("Get remote address error:", err)
	}
	return
}