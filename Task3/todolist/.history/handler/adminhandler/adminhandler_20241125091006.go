package adminhandler

import (
	"context"
	"todolist/model"
	"todolist/task"

	"github.com/cloudwego/hertz/pkg/app"
)

type Params struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var User model.User

func Login(c context.Context, ctx *app.RequestContext) {
	name := ctx.Query("username")
	pwd := ctx.Query("password")
	//验证用户名密码
	err := task.TAsk(name, pwd)
	if err != nil {
		panic(err)
	}
	//验证成功就登录成功
}

func Logout(c context.Context, ctx *app.RequestContext) {

}
