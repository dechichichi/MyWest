package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Data 嵌套结构体，包含文章相关数据
type Data struct {
	Title     string `gorm:"type:varchar(100);uniqueIndex"`
	State     bool
	Content   string `gorm:"type:text"`
	Views     int
	Status    int
	CreatedAt time.Time
	StartTime time.Time
	EndTime   time.Time
}

// User 用户模型，包含用户信息和文章数据
type User struct {
	ID       string `gorm:"primaryKey"`
	Password string `gorm:"type:varchar(50)"`
	Data     Data   `gorm:"embedded"`
}

func main() {
	// 数据库连接字符串，需要根据实际情况替换
	dsn := "username:password@tcp(127.0.0.1:3306)/dbname?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 设置连接池参数
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Hour)

	// 启用日志，方便调试
	db.LogMode(true)

	// 自动迁移
	db.AutoMigrate(&User{})

	// 创建用户
	user := User{
		ID:       "1",
		Password: "password123",
		Data: Data{
			Title:     "GORM 嵌套结构体示例",
			State:     true,
			Content:   "这是一篇关于如何使用 GORM 处理嵌套结构体的文章。",
			Views:     0,
			Status:    1,
			CreatedAt: time.Now(),
			StartTime: time.Now(),
			EndTime:   time.Now().Add(24 * time.Hour),
		},
	}

	// 保存用户
	if err := db.Create(&user).Error; err != nil {
		log.Fatal(err)
	}

	fmt.Println("用户创建成功！")

	// 查询用户
	var foundUser User
	if err := db.Where("id = ?", "1").First(&foundUser).Error; err != nil {
		log.Fatal(err)
	}

	fmt.Printf("找到用户：%+v\n", foundUser)

	// 更新用户
	foundUser.Data.Views = 10
	if err := db.Save(&foundUser).Error; err != nil {
		log.Fatal(err)
	}

	fmt.Println("用户更新成功！")

	// 删除用户
	if err := db.Delete(&foundUser, "1").Error; err != nil {
		log.Fatal(err)
	}

	fmt.Println("用户删除成功！")
}
