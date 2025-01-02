package main

import (
	"log"
	"todolist/config"
	"todolist/routes"
	"todolist/task"
)

func main() {
	// 初始化配置
	config.Init()

	// 初始化数据库连接
	db, err := task.Init(config.Dsn)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 获取数据库实例
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}
	defer sqlDB.Close()

	// 初始化路由
	r := routes.Router()

	// 启动服务器
	if err := r.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
