package dns

import (
	"github.com/SUCHMOKUO/falcon-ws/util"
	"log"
	"net"
	"testing"
)

func TestInsertRecordsToDB(t *testing.T) {
	log.Println("path:", util.GetCurrentPath())
	ip, err := insertRecordToDB("test")
	if err != nil {
		t.Error(err)
	}
	log.Println("insert ip:", ip)
}

func TestGetIPOfDomainFromDB(t *testing.T) {
	log.Println("path:", util.GetCurrentPath())
	expect := net.IPv4(10, 192, 0, 2)
	ip, err := getIPOfDomainFromDB("test")
	if err != nil {
		t.Error(err)
	}
	if !expect.Equal(ip) {
		t.Errorf("should be %v, but get %v\n", expect, ip)
	}
}

func TestWhitelist(t *testing.T) {
	domain := "www.baidu.com"
	isIn := isInWhiteList(domain)
	if isIn {
		t.Error("should not in whitelist")
	}
	insertDomainToDB(domain)
	isIn = isInWhiteList(domain)
	if !isIn {
		t.Error("should in whitelist")
	}
}
