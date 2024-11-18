package routes

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Router() {
	h := server.New(server.WithHostPorts(":8080"))
	h1 := h.Group("/data")
	h2 := h1.Group("/item")
	//中间件
	h.Use()
	//
	h.GET("/login", func(c context.Context, ctx *app.RequestContext) {
		//业务逻辑
	})
	h.GET("/logout", func(c context.Context, ctx *app.RequestContext) {
		//业务逻辑
	})
	h.GET("/register", func(c context.Context, ctx *app.RequestContext) {
		//业务逻辑
	})
	h1.POST("", func(c context.Context, ctx *app.RequestContext) {
		//业务逻辑
	})
	h2.POST("/put/:id", func(c context.Context, ctx *app.RequestContext) {
		//业务逻辑
	})
	h2.GET("/write/:id", func(c context.Context, ctx *app.RequestContext) {
		//业务逻辑
	})
	h.Spin()
}

//http部分由三部分框架组成
