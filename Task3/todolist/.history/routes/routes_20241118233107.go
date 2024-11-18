package routes

import (
	"context"
	"todolist/handler/adminhandler"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Router() {
	h := server.New(server.WithHostPorts(":8080"))
	h1 := h.Group("/admin")
	h2 := h.Group("/item")
	//中间件
	h.Use()
	//
	h1.POST("/login", func(c context.Context, ctx *app.RequestContext) {
		adminhandler.CheckParams(ctx)
		//验证用户名密码
	})

	h.Spin()
}

//http部分由三部分框架组成
