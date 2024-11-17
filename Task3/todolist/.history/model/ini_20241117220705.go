package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Data struct {
	Title     string
	State     bool
	Content   string
	Views     int
	Status    int
	CreatedAt time.Time
	StartTime time.Time
	EndTime   time.Time
}

type User struct {
	ID       string
	Password string
	Data     Data `gorm:"embedded"`
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
	db.AutoMigrate(&User{})
}
