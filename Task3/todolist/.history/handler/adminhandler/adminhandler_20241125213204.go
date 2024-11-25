package adminhandler

import (
	"context"
	"net/http"
	"todolist/model"
	"todolist/task"

	JWT "todolist/pkg/jwt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
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
	claims, _ := JWT.JwtMiddleware.GetClaimsFromJWT(ctx, c)
	token, _ := JWT.JwtMiddleware.ParseToken(ctx, c)
	tokenString, err := token.SignedString("WHDUIAJHDFUAWHDUGAUDAUGDUAIWOWUIUWYHEUGAYFDGAWETYGYIKKKKAISDHAIUDHF")
	if err != nil {
		c.JSON(200, utils.H{
			"message": "登录失败",
			"code":    http.StatusBadRequest,
		})
	}
	c.JSON(200, utils.H{
		"message": "登录成功",
		"code":    http.StatusOK,
		"token":   tokenString,
		"claims":  claims})
	return err
}

func Logout(ctx context.Context, c *app.RequestContext) {
	//清除token
	JWT.JwtMiddleware.LogoutResponse(ctx, c, http.StatusOK)
	c.JSON(200, utils.H{
		"message": "退出成功",
		"code":    http.StatusOK,
	})
}
