package relay

import (
	"net"
	"strconv"
)

// TCPRelayConn is the wrapper of
// net.Conn used for tcp relay.
type TCPRelayConn struct {
	net.Conn
	relay *TCPRelay
}

// Close rewrite the Close method of net.Conn,
// it put the nat port back to nat pool after
// conn closed.
func (conn TCPRelayConn) Close() error {
	// get nat port, and put it into nat pool.
	addr := conn.RemoteAddr().String()
	_, port, _ := net.SplitHostPort(addr)
	portInt, _ := strconv.Atoi(port)
	conn.relay.NAT.DelRecord(uint16(portInt))
	// close the conn.
	return conn.Conn.Close()
}