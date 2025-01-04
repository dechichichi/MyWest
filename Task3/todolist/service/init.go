package service

import (
	"fmt"
	"time"
	"todolist/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var dsn string

func Init() (*gorm.DB, error) {
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 获取底层的 *sql.DB 对象
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 设置连接池参数
	sqlDB.SetMaxOpenConns(100)          // 设置最大打开连接数
	sqlDB.SetMaxIdleConns(50)           // 设置最大空闲连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 设置连接的最大生存时间

	return db, nil
}

func GwtDSN() {
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.ConfigData.DB.Dbuser, config.ConfigData.DB.Dbpassword, config.ConfigData.DB.Dbhost, config.ConfigData.DB.Dbport, config.ConfigData.DB.Dbdatabase)
}
