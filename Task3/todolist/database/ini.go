package database

import (
	"time"
)

// Data 嵌套结构体，包含文章相关数据
type Data struct {
	Title     string `gorm:"type:varchar(100);uniqueIndex"`
	State     bool
	Content   string `gorm:"type:text"`
	Views     int
	Status    int
	CreatedAt time.Time
	StartTime time.Time `gorm:"type:DATETIME;default:CURRENT_TIMESTAMP"`
	EndTime   time.Time `gorm:"type:DATETIME;default:CURRENT_TIMESTAMP"`
}

// User 用户模型，包含用户信息和文章数据
type User struct {
	UserName string `gorm:"type:varchar(50);uniqueIndex"`
	ID       string `gorm:"primaryKey"`
	Password string `gorm:"type:varchar(50)"`
	Data     Data   `gorm:"embedded"`
	Email    string `gorm:"type:varchar(50);uniqueIndex"`
}
