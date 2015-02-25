package pqlap

import (
    _"github.com/lib/pq"
    "database/sql"
    _"fmt"
    "log"
    _"os"
)

var db *sql.DB
var err error

func Sum(x int, y int) int {
    return x + y
}

func connection() (*sql.DB, error) {
    return sql.Open("postgres", "user=iyo password=certate host=db1 dbname=godbtest sslmode=disable")
}

func Close() {
    log.Printf("db close\n")
    db.Close()
}

func ConnectDb() (*sql.DB, error) {
    if db == nil {
        db, err = connection()
    }
    err = db.Ping()
    return db, err
}
