package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
	db, err := sql.Open("mysql", "root:root1234@tcp(127.0.0.1:13306)/db1?charset=utf8")
	checkErr(err)

	stmt, err := db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err := stmt.Exec(1)
	checkErr(err)

	fmt.Println(res.LastInsertId())
}