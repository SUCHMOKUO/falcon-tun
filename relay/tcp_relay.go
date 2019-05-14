package relay

import (
	"github.com/SUCHMOKUO/falcon-tun/dns"
	"github.com/SUCHMOKUO/falcon-tun/nat"
	"io"
	"log"
	"net"
	"strconv"
)

// TCPConnHandler is the type alias of connection handler function.
type TCPConnHandler = func(host, port string, conn io.ReadWriteCloser)

// TCPRelay represent a relay for tcp connection.
type TCPRelay struct {
	NAT *nat.NAT4
	Addr *net.TCPAddr
	HandleConn TCPConnHandler
}

// NewTCPRelay return a new instance of TCPRelay.
func NewTCPRelay(addr *net.TCPAddr) *TCPRelay {
	relay := new(TCPRelay)
	relay.NAT = nat.New()
	relay.Addr = addr
	return relay
}

// Run create a new tcp relay.
func (tr *TCPRelay) Run() {
	l, err := net.ListenTCP("tcp4", tr.Addr)
	if err != nil {
		log.Fatalln("TCP relay error:", err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}
		host, port := tr.getTargetInfo(conn)
		tcpRelayConn := &TCPRelayConn{
			relay: tr,
			Conn: conn,
		}
		go tr.HandleConn(host, port, tcpRelayConn)
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
	_, _, dstPort := tr.NAT.GetRecord(uint16(natPort))
	port = strconv.Itoa(int(dstPort))
	// get target domain.
	domain, err = dns.GetDomain(ipStr)
	if err != nil {
		log.Fatalln("Get remote address error:", err)
	}
	return
}