package model

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

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

func DbInit(consstring string) {
	db, err := gorm.Open(consstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)
	if gin.Mode() == gin.TestMode {
		db.DropTableIfExists(&User{})
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)
	db.AutoMigrate(&User{})
}
