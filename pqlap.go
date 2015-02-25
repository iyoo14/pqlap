package pqlap

import (
    _"github.com/lib/pq"
    "database/sql"
    _"fmt"
    _"os"
)

var instance *sql.DB
var err error

func Sum(x int, y int) int {
    return x + y
}

func connection() (*sql.DB, error) {
    return sql.Open("postgres", "user=iyo password=certate host=db1 dbname=godbtest sslmode=disable")
}

func GetInstance() (*sql.DB, error) {
    if instance == nil {
        instance, err = connection()
    }
    err = instance.Ping()
    return instance, err
}
