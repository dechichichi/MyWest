package routes

import (
	"time"
	"todolist/handler/adminhandler"
	"todolist/handler/taskhandler"
	"todolist/handler/userhandler"
	"todolist/pkg/jwt"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/cache"
	"github.com/hertz-contrib/cache/persist"
	"github.com/hertz-contrib/logger/accesslog"
)

func Router() {
	h := server.New(server.WithHostPorts(":8080"))
	//中间件
	h.Use(accesslog.New())
	memoryStore := persist.NewMemoryStore(1 * time.Minute)
	h.Use(cache.NewCacheByRequestURI(memoryStore, 2*time.Second))
	mw := jwt.MyJwt()
	//admin部分
	h1 := h.Group("todolist/")
	{
		h1.POST("user/register", userhandler.Register)
		h1.POST("user/login", adminhandler.Login)
	}
	auth := h1.Group("auth")
	auth.Use(mw)
	{
		auth.GET("/:id", adminhandler.Login) //test
		auth.POST("/:id/task/create", taskhandler.CreateTask)
		auth.GET("/:id/task/:tid", taskhandler.List)
		auth.POST("/:id/tasks", taskhandler.GetTasksToDone)
		auth.POST("/:id/task/search", taskhandler.GetTasksToKey)
		auth.PUT("/:id/task/:tid", taskhandler.UpdateTask)
		auth.DELETE("/:id/task/:tid", taskhandler.DeleteTask)
	}
}

//http部分由三部分框架组成
