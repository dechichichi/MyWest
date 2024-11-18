package task

import (
	"todolist/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// 增
func Add(username string, password string, email string) (*model.User, error) {

}

// 删
func Delete(username string, password string) error {

}

// 改
func ModifyName(username string, password string, newname string) (*model.User, error) {

}

func ModifyPassword(username string, password string, newpassword string) (*model.User, error) {

}

// 查
func Ask(username string) (*model.User, error) {
	db, err := gorm.Open("mysql", "root:Ly05985481282@/ginclass?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	var user model.User
	//查询数据库
	db.Where("username = ?", username).First(&user)
	//返回user
}
