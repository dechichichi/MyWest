package adminhandler

import (
	"context"
	"todolist/model"

	"github.com/cloudwego/hertz/pkg/app"
)

type Params struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c context.Context, ctx *app.RequestContext) {
	var User model.User
	name := ctx.Query("username")
	pwd := ctx.Query("password")
	err := check(name, pwd)
	//验证用户名密码
}
func check(name, pwd string) error {

	returrn
}

func Logout(c context.Context, ctx *app.RequestContext) {

}
