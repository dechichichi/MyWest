package model

import "github.com/jinzhu/gorm"

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

func Dbdatabase(consstring string) {
	db, err := gorm.Open(consstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&User{})
}
