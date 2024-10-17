package model

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	User      User   `gorm:"foreignkey:UID"`
	UID       uint   `gorm:"not null"`
	Title     string `gorm:"index:not null"`
	Status    int    `gorm:"default:0"` //0:未完成
	Content   string `gorm:"type:longtext"`
	StartTime int64
	EndTime   int64
}
