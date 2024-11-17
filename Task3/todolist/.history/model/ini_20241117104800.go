package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func DbInit(consstring string) {
	db, err := gorm.Open("mysql", "root:Ly05985481282@/todolist?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
}
