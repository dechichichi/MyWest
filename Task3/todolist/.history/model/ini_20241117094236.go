package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

func DbInit(consstring string) {
	db, err := gorm.Open("mysql", consstring)
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)
}
