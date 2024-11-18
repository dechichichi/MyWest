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

	//验证用户名密码
}

func Logout(c context.Context, ctx *app.RequestContext) {

}

func Register(*app.RequestContext) error {

}
