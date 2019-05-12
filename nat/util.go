package nat

import (
	"github.com/SUCHMOKUO/falcon-tun/util"
	"net"
)

// convert record to uint64.
// used by nat function.
func recordToUint64(srcPort uint16, dstIP net.IP, dstPort uint16) uint64 {
	dstIP = dstIP.To4()
	res := uint64(srcPort) << 48
	res |= uint64(util.IPv4ToUint32(dstIP)) << 16
	res |= uint64(dstPort)
	return res
}

// convert uint64 to record.
// used by nat function.
func uint64ToRecord(n uint64) (srcPort uint16, dstIP net.IP, dstPort uint16) {
	dstPort = uint16(n)
	dstIP = util.Uint32ToIPv4(uint32(n >> 16))
	srcPort = uint16(n >> 48)
	return
}
