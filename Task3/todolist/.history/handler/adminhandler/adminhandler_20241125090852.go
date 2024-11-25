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
	err := check(name, pwd)
	//验证用户名密码
}
func check(name, pwd string) error {
	return task.TAsk(name, pwd)

}

func Logout(c context.Context, ctx *app.RequestContext) {

}
