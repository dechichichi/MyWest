package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type Data struct {
	title      string
	state      bool
	content    string
	views      int
	status     int
	created_at string
	start_time string
	end_time   string
}

type User struct {
	ID       string
	Password string
	Data     Data
}

func UserInit(consstring string) {
	db, err := gorm.Open("mysql", consstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	db.AutoMigrate(&User{})
}

func DataInit(consstring string) {
	db, err := gorm.Open("mysql", consstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	db.AutoMigrate(&Data{})
}
