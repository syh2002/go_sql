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

	defer db.Close()

	stmt, err := db.Prepare("update userinfo set username =? where uid =?")
	checkErr(err)

	res, err := stmt.Exec("xh", 1)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
}