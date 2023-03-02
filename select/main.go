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


	// type Stmt interface {
	// 	Close() error
	// 	NumInput() int
	// 	Exec(args []Value) (Result, error)
	// 	Query(args []Value) (Rows, error)
	// }
	// Close函数关闭当前的链接状态，但是如果当前正在执行query，query还是有效返回rows数据。
	
	// NumInput函数返回当前预留参数的个数，当返回>=0时数据库驱动就会智能检查调用者的参数。当数据库驱动包不知道预留参数的时候，返回-1。
	
	// Exec函数执行Prepare准备好的sql，传入参数执行update/insert等操作，返回Result数据
	
	// Query函数执行Prepare准备好的sql，传入需要的参数执行select操作，返回Rows结果集

	rows, err := db.Query("select * from userinfo")
	checkErr(err)

	// Rows是执行查询返回的结果集接口定义


	// type Rows interface {
	// 	Columns() []string
	// 	Close() error
	// 	Next(dest []Value) error
	// }
	// Columns函数返回查询数据库表的字段信息，这个返回的slice和sql查询的字段一一对应，而不是返回整个表的所有字段。

	// Close函数用来关闭Rows迭代器。

	// Next函数用来返回下一条数据，把数据赋值给dest。dest里面的元素必须是driver.Value的值除了string，返回的数据里面所有的string都必须要转换成[]byte。如果最后没数据了，Next函数最后返回io.EOF。

	for rows.Next() {
        var uid int
        var username string
        var department string
        var created string
        err = rows.Scan(&uid, &username, &department, &created)
		// 在scan函数里面如何把driver.Value值转化成用户定义的值
        checkErr(err)
        fmt.Println(uid)
        fmt.Println(username)
        fmt.Println(department)
        fmt.Println(created)
    }
}