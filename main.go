package main

import (
	"github.com/SUCHMOKUO/falcon-tun/setting"
	"github.com/SUCHMOKUO/falcon-tun/tun"
	"github.com/SUCHMOKUO/falcon-ws/client"
	"github.com/SUCHMOKUO/falcon-ws/util"
	"io"
	"log"
)

func main() {
	falconCfg := setting.GetFalconConfig()
	falcon := client.New(falconCfg)
	handleTCPConn := func(host, port string, conn io.ReadWriteCloser) {
		t, err := falcon.CreateTunnel(host, port)
		if err != nil {
			log.Println("dial proxy server error:", err)
			_ = conn.Close()
			return
		}
		go util.Copy(t, conn)
		go util.Copy(conn, t)
	}

	tunIP := setting.GetTUNIP()
	tunCfg := &tun.Config{
		IP: tunIP,
		HandleTCPConn: handleTCPConn,
	}
	tunService := tun.New(tunCfg)
	tunService.Run()
}