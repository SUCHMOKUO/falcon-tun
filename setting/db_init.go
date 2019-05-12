package setting

import (
	"database/sql"
	"github.com/SUCHMOKUO/falcon-ws/util"
	"log"
	"path/filepath"
)

var (
	db *sql.DB
)

func init() {
	prepareDB()
	initTable()
}

func prepareDB() {
	dbPath := filepath.Join(util.GetCurrentPath(), "setting.db")
	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalln("open db error:", err)
	}
}

func initTable() {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS setting (
			server_addr TEXT NOT NULL,
			secure TEXT NOT NULL,
			ipv6 TEXT NOT NULL,
			lookup TEXT NOT NULL,
			fake_host TEXT NOT NULL,
			user_agent TEXT NOT NULL,
			dns_addr TEXT NOT NULL,
			tun_net TEXT NOT NULL
		)
	`)

	if err != nil {
		log.Fatalln("init table error:", err)
	}
}
