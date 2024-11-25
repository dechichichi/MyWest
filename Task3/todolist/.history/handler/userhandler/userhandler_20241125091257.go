package userhandler

import (
	"context"
	"todolist/task"

	"github.com/cloudwego/hertz/pkg/app"
)

func Register(c context.Context, ctx *app.RequestContext) {
	name := ctx.Query("username")
	passward := ctx.Query("passward")
	email := ctx.Query("email")
	err := task.Ask(name)
	if err != nil {
		panic(err)
	}
}
