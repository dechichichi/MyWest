package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username       string `gorm:"unique"`
	passwordDigest string //存储密文
}
