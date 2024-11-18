package routes

import (
	"context"
	"todolist/handler/adminhandler"
	"todolist/handler/taskhandler"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Router() {
	h := server.New(server.WithHostPorts(":8080"))
	h1 := h.Group("/admin")
	h2 := h.Group("/task")
	//中间件
	h.Use()
	//
	h1.POST("/login", func(c context.Context, ctx *app.RequestContext) {
		adminhandler.Check(ctx)
		//验证用户名密码
	})
	h1.POST("/logout", func(c context.Context, ctx *app.RequestContext) {
		adminhandler.Logout(ctx)
	})
	h1.POST("/register", func(c context.Context, ctx *app.RequestContext) {
		userhandler.Register(ctx)
	})
    //item部分
	h2.GET("/list", func(c context.Context, ctx *app.RequestContext) {
		taskhandler.List(ctx)
	})
	h2.POST("/create", func(c context.Context, ctx *app.RequestContext) {
		taskhandler.CreateTask(ctx)
	})
	h2.POST("/update", func(c context.Context, ctx *app.RequestContext) {
		taskhandler.UpdateTask(ctx)
	})
	h2.POST("/delete", func(c context.Context, ctx *app.RequestContext) {
		taskhandler.DeleteTask(ctx)
	h.Spin()
}

//http部分由三部分框架组成
