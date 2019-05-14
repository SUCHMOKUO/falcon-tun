package dns

import (
	"database/sql"
	"github.com/SUCHMOKUO/falcon-ws/util"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"path/filepath"
)

func init() {
	prepareDB()
	initTable()
	prepareStmt()
}

var (
	db              *sql.DB
	insertWhitelist *sql.Stmt
	insertRecords   *sql.Stmt
)

func prepareDB() {
	dbPath := filepath.Join(util.GetCurrentPath(), "dns.db")
	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalln("open db error:", err)
	}
}

func initTable() {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS records (
			ip INTEGER PRIMARY KEY AUTOINCREMENT,
			domain TEXT NOT NULL
		);
		
		CREATE INDEX IF NOT EXISTS domain_index ON records(domain);
		
		CREATE TABLE IF NOT EXISTS whitelist (
			domain TEXT PRIMARY KEY
		);
		
		CREATE TABLE IF NOT EXISTS whitelist_tmp (
			domain TEXT PRIMARY KEY
		);
	`)

	if err != nil {
		log.Fatalln("init table error:", err)
	}

	db.Exec(`
		INSERT INTO records(ip, domain) VALUES(1, "0");
		DELETE FROM records WHERE ip = 1;
	`)
}

func prepareStmt() {
	var err error
	insertRecords, err = db.Prepare(`INSERT INTO records(domain) VALUES(?)`)
	if err != nil {
		log.Fatalln("prepare insert_records stmt error:", err)
	}
	insertWhitelist, err = db.Prepare(`INSERT INTO whitelist(domain) VALUES(?)`)
	if err != nil {
		log.Fatalln("prepare insert_whitelist stmt error:", err)
	}
}
