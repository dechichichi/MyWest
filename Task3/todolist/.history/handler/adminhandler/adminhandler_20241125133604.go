package adminhandler

import (
	"context"
	"todolist/model"
	"todolist/task"

	JWT "todolist/pkg/jwt"

	"github.com/cloudwego/hertz/pkg/app"
)

type Params struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var User model.User

func Login(ctx context.Context, c *app.RequestContext) {
	name := c.Query("username")
	pwd := c.Query("password")
	//验证用户名密码
	err := task.TAsk(name, pwd)
	if err != nil {
		panic(err)
	}
	//验证成功就登录成功
	//接下来生成token返回给前端
	claims, err := JWT.JwtMiddleware.GetClaimsFromJWT(ctx, c)
	token, err := JWT.JwtMiddleware.ParseToken(ctx, c)
}

func Logout(ctx context.Context, c *app.RequestContext) {

}
