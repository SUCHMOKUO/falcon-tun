package nat

import (
	"log"
	"net"
	"testing"
)

func TestRecordToUint64(t *testing.T) {
	tests := []struct{
		srcPort uint16
		dstIP net.IP
		dstPort uint16
		res uint64
	} {
		{ 0, net.IPv4(0, 0, 0, 0), 0, 0 },
		{ 40000, net.IPv4(10, 192, 50, 1), 443, 11259010889015165371 },
	}

	for _, test := range tests {
		if res := recordToUint64(test.srcPort, test.dstIP, test.dstPort); res != test.res {
			t.Errorf("toUint64: %v, expect %d, but get %d\n", test, test.res, res)
		}
	}
}

func TestUint64ToRecord(t *testing.T) {
	tests := []struct{
		n uint64
		srcPort uint16
		dstIP net.IP
		dstPort uint16
	} {
		{ 0, 0, net.IPv4(0, 0, 0, 0), 0 },
		{ 11259010889015165371, 40000, net.IPv4(10, 192, 50, 1), 443 },
	}

	for _, test := range tests {
		srcPort, dstIP, dstPort := uint64ToRecord(test.n)
		if srcPort != test.srcPort || dstPort != test.dstPort || !dstIP.Equal(test.dstIP) {
			t.Errorf("uint64ToRecord: %v\nget srcPort: %d, dstIP: %v, dstPort: %d\n", test, srcPort, dstIP, dstPort)
		}
	}
}

func TestNAT4_AddRecord(t *testing.T) {
	nat4 := New()
	natPort := nat4.AddRecord(40000, net.IPv4(10, 192, 50, 1), 443)
	log.Println(natPort)
}

func TestNAT4_GetRecord(t *testing.T) {
	nat4 := New()
	var srcPort, dstPort uint16 = 40000, 443
	dstIP := net.IPv4(10, 192, 50, 1)
	natPort := nat4.AddRecord(srcPort, dstIP, dstPort)
	getSrcPort, getDstIP, getDstPort := nat4.GetRecord(natPort)
	if !dstIP.Equal(getDstIP) || srcPort != getSrcPort || dstPort != getDstPort {
		t.Errorf("%d, %v, %d => %d, %v, %d\n", srcPort, dstIP, dstPort, getSrcPort, getDstIP, getDstPort)
	}
}

func TestNAT4_DelRecord(t *testing.T) {
	nat4 := New()
	l := len(nat4.portPool)
	var srcPort, dstPort uint16 = 40000, 443
	dstIP := net.IPv4(10, 192, 50, 1)
	natPort := nat4.AddRecord(srcPort, dstIP, dstPort)
	if len(nat4.portPool) != l-1 {
		t.Errorf("before put: len: %d, original: %d\n", len(nat4.portPool), l)
	}
	nat4.DelRecord(natPort)
	if len(nat4.portPool) != l {
		t.Errorf("after put: len: %d, original: %d\n", len(nat4.portPool), l)
	}
	_, ip, _ := nat4.GetRecord(natPort)
	if ip != nil {
		t.Error(ip)
	}
}