package routes

import (
	"todolist/handler/adminhandler"
	"todolist/handler/taskhandler"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func Router() {
	h := server.New(server.WithHostPorts(":8080"))
	h1 := h.Group("/admin")
	h2 := h.Group("/task")
	//中间件
	h.Use()
	//
	h1.POST("/login", adminhandler.Login)
	h1.POST("/logout", adminhandler.Logout)
	h1.POST("/register", adminhandler.Register)
	//item部分
	h2.GET("/list", taskhandler.List)
	h2.POST("/create", taskhandler.CreateTask)
	h2.POST("/update", taskhandler.UpdateTask)
	h2.POST("/delete", taskhandler.DeleteTask)
	h.Spin()
}

//http部分由三部分框架组成
