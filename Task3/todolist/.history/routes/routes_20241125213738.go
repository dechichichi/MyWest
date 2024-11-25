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
	// 中间件
	h.Use(accesslog.New())
	memoryStore := persist.NewMemoryStore(1 * time.Minute)
	h.Use(cache.NewCacheByRequestURI(memoryStore, 2*time.Second))
	mw := jwt.MyJwt()

	// 用户注册和登录不需要JWT保护
	h1 := h.Group("/user")
	h1.POST("register", userhandler.Register)
	h1.POST("login", adminhandler.Login)

	// 需要JWT保护的路由
	auth := h1.Group("auth")
	auth.Use(mw)
	{
		auth.GET("/:id", userhandler.Auth) // 确保这是正确的处理函数
		auth.POST("/:id/task/create", taskhandler.CreateTask)
		auth.GET("/:id/task/:tid", taskhandler.List)
		auth.POST("/:id/tasks", taskhandler.GetTasksToDone)
		auth.POST("/:id/task/search", taskhandler.GetTasksToKey)
		auth.PUT("/:id/task/:tid", taskhandler.UpdateTask)
		auth.DELETE("/:id/task/:tid", taskhandler.DeleteTask)
	}

	h.Spin()
}
