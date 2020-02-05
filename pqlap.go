package pqlap

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"sync"
)

var db *Db
var once sync.Once
var err error

type Db struct {
	con    *sql.DB
	txn    *sql.Tx
	stmt   *sql.Stmt
	err    error
	Result sql.Result
}

func DbConnection(dsn string) *Db {
	once.Do(func() {
		fmt.Println("db connection.")
		d, err := sql.Open("postgres", dsn)
		db = &Db{}
		db.err = err
		if err == nil {
			db.con = d
			db.err = d.Ping()
		}
	})
	return db
}

func (d *Db) Error() bool {
	if d.err != nil {
		return true
	}
	return false
}

func (d *Db) GetError() error {
	return d.err
}

func (d *Db) Begin() {
	txn, err := d.con.Begin()
	d.txn = txn
	d.err = err
}

func (d *Db) Commit() {
	err := d.txn.Commit()
	d.err = err
}

func (d *Db) Rollback() {
	err := d.txn.Rollback()
	d.err = err
}

func (d *Db) Close() {
	err := d.con.Close()
	d.err = err
}

func (d *Db) Prepare(sql string) {
	stmt, err := d.txn.Prepare(sql)
	d.stmt = stmt
	d.err = err
}

func (d *Db) Exec(val []interface{}) {
	result, err := d.stmt.Exec(val...)
	d.Result = result
	d.err = err
}
