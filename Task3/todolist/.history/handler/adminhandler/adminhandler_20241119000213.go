package adminhandler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

type Params struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c context.Context, ctx *app.RequestContext) {
     name:=ctx.Query("username")
	 pwd:=ctx.Query("password")
     err:=check(name,pwd)
	//验证用户名密码
}
func check(name,pwd string) error{

	returrn nil
}

func Logout(c context.Context, ctx *app.RequestContext) {

}

func Register(c context.Context, ctx *app.RequestContext) {

}
