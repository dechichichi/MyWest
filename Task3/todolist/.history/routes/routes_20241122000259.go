package routes

import (
	"time"
	"todolist/handler/adminhandler"
	"todolist/handler/taskhandler"
	"todolist/handler/userhandler"
	"todolist/pkg/jwt"
	"todolist/pkg/keyauth"
	"todolist/pkg/sessions"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/cache"
	"github.com/hertz-contrib/cache/persist"
)

func Router() {
	h := server.New(server.WithHostPorts(":8080"))
	h1 := h.Group("/admin")
	h2 := h.Group("/task")
	//中间件
	h.Use(sessions.MySession())
	h.Use(jwt.MyJwt())
	h.Use(keyauth.MyKeyauth())
	memoryStore := persist.NewMemoryStore(1 * time.Minute)
	h.Use(cache.NewCacheByRequestURI(memoryStore, 2*time.Second))
	h1.Use(jwt.MyJwt())
	//
	h1.POST("/login", adminhandler.Login)
	h1.POST("/logout", adminhandler.Logout)
	h1.POST("/register", userhandler.Register)
	//item部分
	h2.GET("/list", taskhandler.List)
	h2.POST("/create", taskhandler.CreateTask)
	h2.POST("/update", taskhandler.UpdateTask)
	h2.POST("/delete", taskhandler.DeleteTask)
	h.Spin()
}

//http部分由三部分框架组成
