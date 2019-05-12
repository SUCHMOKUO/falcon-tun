package setting

import (
	"github.com/SUCHMOKUO/falcon-ws/client"
	"log"
	"net"
)

// GetFalconConfig return a falcon config
// get from the database.
func GetFalconConfig() *client.Config {
	cfg := getConfig()
	falconCfg := new(client.Config)
	falconCfg.ServerAddr = cfg.serverAddr
	falconCfg.FakeHost = cfg.fakeHost
	falconCfg.UserAgent = cfg.userAgent
	falconCfg.Secure = getBool(cfg.secure)
	falconCfg.IPv6 = getBool(cfg.ipv6)
	falconCfg.Lookup = getBool(cfg.lookup)
	return falconCfg
}

// GetDNSAddr return the dns server
// listen address from database.
func GetDNSAddr() string {
	cfg := getConfig()
	return cfg.dnsAddr
}

// GetTUNNet return the tun device
// network address from database.
func GetTUNNet() string {
	cfg := getConfig()
	return cfg.tunNet
}

// GetTUNNetIP return the net ip of tun net.
func GetTUNNetIP() net.IP {
	tunNet := GetTUNNet()
	_, ipNet, err := net.ParseCIDR(tunNet)
	if err != nil {
		log.Fatalln("tun_net format error:", err)
	}
	return ipNet.IP
}