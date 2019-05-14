package dns

import (
	"net"
	"testing"
)

func TestGetFakeIP(t *testing.T) {
	ip, err := GetFakeIP("test")
	if err != nil {
		t.Error(err)
	}
	expect := net.IPv4(10, 192, 0, 2)
	if !expect.Equal(ip) {
		t.Errorf("should be %v, but get %s\n", expect, ip)
	}
}

func TestGetDomain(t *testing.T) {
	domain, err := GetDomain("10.192.0.2")
	if err != nil {
		t.Error(err)
	}
	if domain != "test" {
		t.Errorf("should be 'test', but get %s\n", domain)
	}
}