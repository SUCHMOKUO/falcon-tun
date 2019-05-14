package dns

import (
	"errors"
	"github.com/SUCHMOKUO/falcon-ws/util"
	_ "github.com/mattn/go-sqlite3"
	"net"
)

var errNotSupportIPv6 = errors.New("not support ipv6 yet.")

// GetDomain return the host of the fake ip.
func GetDomain(ipStr string) (string, error) {
	ip := net.ParseIP(ipStr)
	if util.IsIPv6(ip) {
		return "", errNotSupportIPv6
	}
	return getDomainOfIPFromDB(ip)
}

// GetFakeIP return the fake ip of domain.
func GetFakeIP(domain string) (net.IP, error) {
	ip, err := getIPOfDomainFromDB(domain)
	if err != nil {
		return nil, err
	}
	return ip, nil
}
