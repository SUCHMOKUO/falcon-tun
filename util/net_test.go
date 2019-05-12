package util

import (
	"net"
	"testing"
)

func TestIPv4ToUint32(t *testing.T) {
	tests := []struct{
		ip net.IP
		res uint32
	} {
		{ net.IPv4(0, 0, 0, 0), 0 },
		{ net.IPv4(1, 1, 1, 1), 16843009 },
		{ net.IPv4(10, 192, 50, 1), 180367873 },
	}

	for _, test := range tests {
		if res := IPv4ToUint32(test.ip); res != test.res {
			t.Error("ipv4ToUint32 error:", test)
		}
	}
}

func TestUint32ToIPv4(t *testing.T) {
	tests := []struct{
		n uint32
		ip net.IP
	} {
		{ 0, net.IPv4(0, 0, 0, 0) },
		{ 16843009, net.IPv4(1, 1, 1, 1) },
		{ 180367873, net.IPv4(10, 192, 50, 1) },
	}

	for _, test := range tests {
		if ip := Uint32ToIPv4(test.n); !test.ip.Equal(ip) {
			t.Errorf("uint32ToIPv4: %v\nexpect %v, get %v\n", test, test.ip, ip)
		}
	}
}