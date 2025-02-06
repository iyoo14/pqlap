package pqlap

import (
	"encoding/json"
	"fmt"
	"log"
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
	con := DbConnection(dns)
	if con.Error() {
		fmt.Println("con error.")
	}
	con.Begin()
	if con.Error() {
		fmt.Println("begin error.")
	}
	con.PrepareTxn("insert into users (id, name, age) values ($1, $2, $3)")
	if con.Error() {
		fmt.Println("prepare error1.")
	}
	var record []interface{}
	id := 1
	record = append(record, id)
	record = append(record, "aaaxxzzxxx")
	record = append(record, 6)
	con.Exec(record)
	if con.Error() {
		fmt.Println("exec error2.")
	}

	st := 13
	con.PrepareTxn("update cc  set state = $1 where stock_code = $2")
	if con.Error() {
		fmt.Println("prepare error3.")
		fmt.Println(con.GetError())
	}
	var record2 []interface{}
	record = append(record2, st)
	record = append(record2, "0000")
	con.Exec(record2)
	if con.Error() {
		fmt.Println("exec error4.")
		fmt.Println(con.GetError())
		con.Rollback()
	}
	con.Commit()

	if con.Error() {
		fmt.Println("commit error.")
	}
}

func TestDbInstantConnection(t *testing.T) {
	dns := getDsn()
	log.Println(dns)
	con := DbConnection(dns)
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
