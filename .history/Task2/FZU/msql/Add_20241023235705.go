package msql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func LinkDB() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/mywest")
	if err != nil {
		return nil, err
	}
	defer db.Close()
}
