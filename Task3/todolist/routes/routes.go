package routes

import (
	"time"
	"todolist/handler/taskhandler"
	"todolist/handler/userhandler"
	"todolist/middleware/jwt"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/cache"
	"github.com/hertz-contrib/cache/persist"
	"github.com/hertz-contrib/logger/accesslog"
)

func Router() *server.Hertz {
	//web服务地址在localhost:8080
	h := server.New(server.WithHostPorts(":8080"))
	// 中间件
	//日志采用默认打印方式 打印在控制台
	h.Use(accesslog.New())
	//设置缓存 有助于提高服务器的并发访问能力
	//设置全局缓存过期时间为1分钟
	memoryStore := persist.NewMemoryStore(1 * time.Minute)
	// 设置针对以 URI 为 Key 的缓存过期时间为2秒
	h.Use(cache.NewCacheByRequestURI(memoryStore, 2*time.Second))

	// 注册API
	// 用户注册和登录不需要JWT保护
	h1 := h.Group("/user")
	h1.POST("register", userhandler.Register)
	h1.POST("login", userhandler.Login)
	// 需要JWT保护的路由
	auth := h1.Group("auth")
	auth.Use(jwt.MyJwt())
	{
		auth.GET("/:id", userhandler.Auth)
		auth.POST("/:id/task/create", taskhandler.CreateTask)
		auth.GET("/:id/task/:tid", taskhandler.List)
		auth.POST("/:id/tasks", taskhandler.GetTasksToDone)
		auth.POST("/:id/task/search", taskhandler.GetTasksToKey)
		auth.PUT("/:id/task/:tid", taskhandler.UpdateTask)
		auth.DELETE("/:id/task/:tid", taskhandler.DeleteTask)
	}
	// 启动服务
	//在接受到关闭请求后,等待所有请求处理完毕，再关闭服务
	h.Spin()
	return h
}
