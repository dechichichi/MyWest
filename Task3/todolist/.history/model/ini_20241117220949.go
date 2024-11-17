package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Data 嵌套结构体，包含文章相关数据
type Data struct {
	Title     string `gorm:"type:varchar(100);uniqueIndex"`
	State     bool
	Content   string `gorm:"type:text"`
	Views     int
	Status    int
	CreatedAt time.Time
	StartTime time.Time
	EndTime   time.Time
}

// User 用户模型，包含用户信息和文章数据
type User struct {
	ID       string `gorm:"primaryKey"`
	Password string `gorm:"type:varchar(50)"`
	Data     Data   `gorm:"embedded"`
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
