package task

import (
	"errors"
	"todolist/model"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Add 添加用户
func Add(username string, password string, email string) (*model.User, error) {
	db, err := gorm.Open(mysql.Open("root:Ly05985481282@/ginclass?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// 检查用户名是否已存在
	var user model.User
	if err := db.Where("username = ?", username).First(&user).Error; err == nil {
		return nil, errors.New("username already exists")
	}

	// 添加用户
	newUser := model.User{ID: username, Password: password, Email: email}
	if err := db.Create(&newUser).Error; err != nil {
		return nil, err
	}

	return &newUser, nil
}

// Delete 删除用户
func Delete(username string, password string) error {
	db, err := gorm.Open(mysql.Open("root:Ly05985481282@/ginclass?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return err
	}

	// 验证用户名和密码
	var user model.User
	if err := db.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return err
	}

	// 删除用户
	if err := db.Delete(&model.User{}, "username = ?", username).Error; err != nil {
		return err
	}

	return nil
}

// ModifyName 修改用户名
func ModifyName(username string, password string, newname string) (*model.User, error) {
	db, err := gorm.Open(mysql.Open("root:Ly05985481282@/ginclass?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 验证用户名和密码
	var user model.User
	if err := db.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return nil, err
	}

	// 修改用户名
	if err := db.Model(&user).Update("Username", newname).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// ModifyPassword 修改密码
func ModifyPassword(username string, password string, newpassword string) (*model.User, error) {
	db, err := gorm.Open(mysql.Open("root:Ly05985481282@/ginclass?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 验证用户名和密码
	var user model.User
	if err := db.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return nil, err
	}

	// 修改密码
	if err := db.Model(&user).Update("Password", newpassword).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Ask 查询用户
func Ask(username string) (*model.User, error) {
	db, err := gorm.Open(mysql.Open("root:Ly05985481282@/ginclass?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	var user model.User
	// 查询数据库
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	// 返回user
	return &user, nil
}
