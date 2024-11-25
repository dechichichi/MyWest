package task

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func Init() {
	db, err = gorm.Open("mysql", "root:Ly05985481282@/ginclass?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
