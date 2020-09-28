package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/sys")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
	//defer Db.Close() // 注意这行代码要写在上面err判断的下面

	// fmt.Println(err)

	// if err != nil {
	// 	fmt.Println("error")
	// 	log.Panicln("err:", err.Error())
	// }
	// fmt.Println("zzz")
	// Db.SetMaxOpenConns(10)
	// Db.SetMaxIdleConns(10)
}
