package task

import (
	"github.com/jinzhu/gorm"
	"gorm.io/driver/mysql"
)

var db *gorm.DB
var err error

func Init() {
	db, err = gorm.Open(mysql.Open("root:Ly05985481282@/ginclass?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
