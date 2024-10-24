package msql

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	consChar := "1"
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
