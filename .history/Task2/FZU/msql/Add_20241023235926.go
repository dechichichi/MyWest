package msql

import (
	"database/sql"

	"gorm.io/gorm"
)

type Info struct {
	Title   string `gorm:"index;not null"`
	Content string `gorm:"type:longtext"`
	Author  string
	Date    string
	Views   string
}
type InfoDao struct {
	*gorm.DB
}

func LinkDB() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/mywest")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return db, nil
}
