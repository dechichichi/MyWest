package userhandler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

func Register(c context.Context, ctx *app.RequestContext) {
	name := ctx.Query("username")
	pwd := ctx.Query("password")
	err := check(name, pwd)
	//验证用户名密码
}
