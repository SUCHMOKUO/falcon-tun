package nat

import (
	"github.com/SUCHMOKUO/falcon-tun/base/bimap"
	"github.com/SUCHMOKUO/falcon-tun/base/queue"
	"log"
	"net"
)

// NAT4 represent a network address translation
// object for IPv4.
type NAT4 struct {
	portPool queue.Queue
	records *bimap.BiMap
}

// New return a new instance of NAT4.
func New() *NAT4 {
	nat4 := new(NAT4)
	portPool := queue.New()
	var port uint16
	for port = 49152; port < 65535; port++ {
		portPool.Put(port)
	}
	nat4.portPool = portPool
	nat4.records = bimap.New()
	return nat4
}

// Record a network address translation for IPv4.
// it returns the translated port.
func (nat4 *NAT4) AddRecord(srcPort uint16, dstIP net.IP, dstPort uint16) uint16 {
	natPort := nat4.portPool.Poll()
	n := recordToUint64(srcPort, dstIP, dstPort)
	err := nat4.records.Put(n, natPort)
	if err != nil {
		log.Fatalln(err)
	}
	return natPort
}

// GetRecord return the record of the nat port.
func (nat4 *NAT4) GetRecord(natPort uint16) (srcPort uint16, dstIP net.IP, dstPort uint16) {
	n, ok := nat4.records.Get(natPort)
	if !ok {
		return
	}
	return uint64ToRecord(n.(uint64))
}

// DelRecord delete a record, and put back
// the nat port into port pool.
func (nat4 *NAT4) DelRecord(natPort uint16) {
	nat4.records.Del(natPort)
	nat4.portPool.Put(natPort)
}