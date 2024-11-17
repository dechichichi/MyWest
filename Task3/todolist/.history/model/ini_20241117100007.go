package model

import (
	"time"

	"github.com/jinzhu/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func DbInit(consstring string) {
	db, err := gorm.Open(mysql.Open(consstring), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
}
