package setting

import (
	"log"
)

type config struct {
	serverAddr string
	secure string
	ipv6 string
	lookup string
	fakeHost string
	userAgent string
	dnsAddr string
	tunNet string
}

var configCache *config

func init() {
	configCache = getConfigFromDB()
}

func getConfigFromDB() *config {
	cfg := new(config)
	q := db.QueryRow(`SELECT * FROM setting`)
	err := q.Scan(
		&cfg.serverAddr,
		&cfg.secure,
		&cfg.ipv6,
		&cfg.lookup,
		&cfg.fakeHost,
		&cfg.userAgent,
		&cfg.dnsAddr,
		&cfg.tunNet)
	if err != nil {
		log.Fatalln("get config error:", err)
	}
	return cfg
}

func getConfig() *config {
	return configCache
}
