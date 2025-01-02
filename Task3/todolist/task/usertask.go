package task

import (
	"errors"
	"fmt"
	"todolist/model"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

// VerifyUser 验证用户名和密码，并返回用户信息
func VerifyUser(username string, password string) (*model.User, error) {
	var user model.User
	// 查询数据库，根据用户名查找用户
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		// 如果没有找到用户，返回错误
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("用户名不存在")
		}
		// 如果是其他类型的错误，直接返回
		return nil, err
	}
	// 比较密码
	if user.Password != password {
		return nil, fmt.Errorf("密码错误")
	}
	// 返回用户信息
	return &user, nil
}

// Add 添加用户
func Add(username string, password string, email string) (*model.User, error) {
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
	// 查询数据库
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	// 返回user
	return &user, nil
}

func TAsk(username string, password string) error {
	_, err := VerifyUser(username, password)
	return err
}

func Auth(username string, password string) (*model.User, error) {
	return VerifyUser(username, password)
}
