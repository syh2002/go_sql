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

	// 插入数据

	// db.Prepare()函数用来返回准备要执行的sql操作，然后返回准备完毕的执行状态。
	// Prepare(query string) (Stmt, error)
	// type Stmt interface {
	// 	Close() error
	// 	NumInput() int
	// 	Exec(args []Value) (Result, error)
	// 	Query(args []Value) (Rows, error)
	// }
	// Exec函数执行Prepare准备好的sql，传入参数执行update/insert等操作，返回Result数据
	stmt, err := db.Prepare("insert into userinfo set username=?,departname=?,created=?")
	// 我们可以看到我们传入的参数都是=?对应的数据，这样做的方式可以一定程度上防止SQL注入。
	checkErr(err)
	// stmt.Exec()函数用来执行stmt准备好的SQL语句
	// type Result interface {
	// 	LastInsertId() (int64, error)
	// 	RowsAffected() (int64, error)
	// }
	// LastInsertId函数返回由数据库执行插入操作得到的自增ID号。
	
	// RowsAffected函数返回query操作影响的数据条目数。
	res, err := stmt.Exec("syh", "研发部门", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
    checkErr(err)

    fmt.Println(id)

	defer db.Close()
}