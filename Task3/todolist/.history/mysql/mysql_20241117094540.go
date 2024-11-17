package mysql

import (
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

func init() {
	db, err := gorm.Open("mysql", "localhost:Ly05985481282@/todolist?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&User{})
}
