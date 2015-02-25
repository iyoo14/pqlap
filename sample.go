package main

import (
    "github.com/lib/pq"
    "database/sql"
    "fmt"
    "os"
    )
type Record struct {
        id int
        displayName sql.NullString
        sex string
        birthday pq.NullTime
        age sql.NullInt64
        married sql.NullBool
        rate sql.NullFloat64
        salary sql.NullInt64
}

func openConnection() (*sql.DB, error) {
    return sql.Open("postgres", "user=iyo password=certate host=db1 dbname=godbtest sslmode=disable")
}

func main() {
    db, err := openConnection()
    checkErr(err)

    err = db.Ping()
    checkErr(err)

    defer db.Close()
    query := "select id, display_name, sex, birthday, age, married, rate, salary from table1 where id >= $1"
    stmt, err := db.Prepare(query)
    checkErr(err)

    rows, err := stmt.Query(1)
    checkErr(err)

    for rows.Next() {
        var r *Record = new(Record)
        recs := []interface{}{&r.id, &r.displayName, &r.sex, &r.birthday, &r.age, &r.married, &r.rate, &r.salary}
        //if err = rows.Scan(&r.id, &r.displayName, &r.sex, &r.birthday, &r.age, &r.married, &r.rate, &r.salary); err != nil {
        if err = rows.Scan(recs...); err != nil {
            fmt.Printf("error rows: %v\n", err)
        }
        fmt.Printf("%d, %+v, %+v, %+v, %+v, %+v, %+v, %+v\n", 
                r.id, r.displayName, r.sex, r.birthday, r.age, r.married, r.rate, r.salary)
        fmt.Printf("%v\n", r.displayName.String)
    }
    rows.Close()
/*
    txn, err := db.Begin()
    checkDbErr(txn, err, "txn")
    //stmt, err = txn.Prepare("INSERT INTO table1 (display_name, sex, birthday, age, married, rate, salary) VALUES ($1, $2, $3, $4, $5, $6, $7)")
    stmt, err = txn.Prepare("INSERT INTO table1 (sex) VALUES ($1)")
    //_, err = stmt.Exec("だんべい", "m", "2014-02-05 10:00:00", 23, 1, 100, 1000)
    _, err = stmt.Exec("m")
    checkDbErr(txn, err, "stmt")

    stmt, err = txn.Prepare("UPDATE table1 SET age = $1 WHERE id = $2")
    checkDbErr(txn, err, "stmt")
    res, err := stmt.Exec(123,1)
    checkDbErr(txn, err, "exec")
    affect, err := res.RowsAffected()
    checkDbErr(txn, err, "rowsaff")
    fmt.Println(affect)
    txn.Commit()
*/
}

func find() {

}
func checkErr(err error) {
    if err != nil {
        fmt.Printf("error: %v\n", err)
        panic(err)
        os.Exit(1)
    }
}

func checkDbErr(txn *sql.Tx, err error, msg string) {
    if  err != nil {
        fmt.Printf("dberr: %v\n", msg)
        txn.Rollback()
        checkErr(err)
    }
}
