package routes

import (
	"time"
	"todolist/handler/adminhandler"
	"todolist/handler/userhandler"
	"todolist/pkg/jwt"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/cache"
	"github.com/hertz-contrib/cache/persist"
	"github.com/hertz-contrib/logger/accesslog"
	"github.com/hertz-contrib/swagger"
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
		auth.GET("/:id", userHandler.IdLogin) //test
		auth.POST("/:id/task/create", taskHandler.CreateTask)
		auth.GET("/:id/task/:tid", taskHandler.CheckTask)
		auth.POST("/:id/tasks", taskHandler.ListTask)
		auth.POST("/:id/task/search", taskHandler.SearchTask)
		auth.PUT("/:id/task/:tid", taskHandler.UpdateTask)
		auth.DELETE("/:id/task/:tid", taskHandler.DeleteTask)
	}

	admin := h1.Group("admin")
	admin.Use(mw)
	{
		admin.POST("/listusers", adminHandler.ListUsers)
		admin.GET("/add/:id", adminHandler.AddAdmin)
		admin.GET("/block/:id", adminHandler.BlockUser)
	}
	url := swagger.URL("http://localhost:8000/swagger/doc.json") // The url pointing to API definition
	h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler, url))
}

//http部分由三部分框架组成
