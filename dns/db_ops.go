package dns

import (
	"database/sql"
	"github.com/SUCHMOKUO/falcon-tun/setting"
	"github.com/SUCHMOKUO/falcon-tun/util"
	"log"
	"net"
)

func isInWhiteList(domain string) bool {
	q := db.QueryRow(`SELECT COUNT(*) FROM whitelist_tmp WHERE domain = ?`, domain)
	count := 0
	err := q.Scan(&count)
	if err == nil && count > 0 {
		return true
	}
	if err != nil {
		panic(err)
	}
	q = db.QueryRow(`SELECT COUNT(*) FROM whitelist WHERE domain = ?`, domain)
	err = q.Scan(&count)
	if err == nil && count > 0 {
		return true
	}
	if err != nil {
		panic(err)
	}
	return false
}

var ipBase uint32

func init() {
	tunNetIP := setting.GetTUNNetIP()
	ipBase = util.IPv4ToUint32(tunNetIP)
}

func getIPOfDomainFromDB(domain string) (ip net.IP, err error) {
	q := db.QueryRow(`SELECT ip FROM records WHERE domain = ?`, domain)
	var ipNum uint16
	err = q.Scan(&ipNum)
	if err != nil {
		log.Println("query ip records error:", err)
		return
	}
	ip = util.Uint32ToIPv4(ipBase + uint32(ipNum))
	return
}

func getDomainOfIPFromDB(ip net.IP) (domain string, err error) {
	id := util.IPv4ToUint32(ip) - ipBase
	q := db.QueryRow(`SELECT domain FROM records WHERE ip = ?`, id)
	err = q.Scan(&domain)
	if err != nil {
		log.Println("query domain records error:", err)
		return
	}
	return
}

// insert domain into database,
// and return the inserted row's ip number.
func insertRecordToDB(domain string) (ip uint16, err error) {
	var res sql.Result
	res, err = insertRecords.Exec(domain)
	if err != nil {
		log.Println("insert records error:", err)
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Println("insert records error:", err)
		return
	}
	ip = uint16(id)
	return
}

// insert domain into whitelist.
func insertDomainToDB(domain string) {
	_, err := insertWhitelist.Exec(domain)
	if err != nil {
		log.Println("insert whitelist error:", err)
	}
}