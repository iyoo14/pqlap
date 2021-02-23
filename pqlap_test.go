package pqlap

import (
	"encoding/json"
	"fmt"
	"log"
	"mywork/db/pqlap"
	"os"
	"path/filepath"
	"testing"
)

type config struct {
	Dsn string `json:"dsn"`
}

func TestConnectDb(t *testing.T) {
	dns := getDsn()
	log.Println(dns)
	con := pqlap.DbConnection(dns)
	if con.Error() {
		fmt.Println("con error.")
	}
	con.Begin()
	if con.Error() {
		fmt.Println("begin error.")
	}
	con.Prepare("insert into test_tbl (id, name) values ($1, $2), ($3, $4)")
	if con.Error() {
		fmt.Println("prepare error.")
	}
	var record []interface{}
	record = append(record, 1)
	record = append(record, "a")
	record = append(record, 2)
	record = append(record, "b")
	con.Exec(record)
	if con.Error() {
		fmt.Println("exec error.")
	}
	con.Commit()
	if con.Error() {
		fmt.Println("commit error.")
	}
}

func TestDbInstantConnection(t *testing.T) {
	dns := getDsn()
	log.Println(dns)
	con := pqlap.DbConnection(dns)
	if con.Error() {
		fmt.Println("con error.")
	}

}

func getDsn() string {
	home := os.Getenv("GOPATH")
	fname := filepath.Join(home, "config", "env.json")
	log.Println(fname)
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal("open error.")
	}
	defer f.Close()
	var cfg config
	err = json.NewDecoder(f).Decode(&cfg)
	return cfg.Dsn
}
