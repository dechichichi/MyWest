package msql

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	consChar := "root:123456@tcp(127.0.0.1:3306)/test"
	//Open函数并不会连接数据库 甚至不会验证其参数 只是把后续struct设置好了
	var err error
	db, err = sql.Open("mysql", consChar)
	if err != nil {

	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
	}

}

func Add(Opeo string) {
	ret, err := db.Exec(Opeo)
	if err != nil {
		panic(err)
	}
	id, err := ret.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("Insert id is:", id)
}

func Query() {
	sqlStr := "select writer,title, content,date from user where id=?"
	var u Thing
	err := db.QueryRow(sqlStr, 1).Scan(&u.Writer, &u.Title, &u.Date, &u.Content)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("writer:%s, title:%s, content:%s, date:%s\n", u.Writer, u.Title, u.Content, u.Date)
}
