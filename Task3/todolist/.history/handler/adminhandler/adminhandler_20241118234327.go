package adminhandler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

type Params struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Check(c context.Context, ctx *app.RequestContext) {

	//验证用户名密码
}

func Adminhandler(*app.HandlerFunc) {

}

func Logout(*app.RequestContext) error {

}
