package adminhandler

import (
	"context"
	"net/http"
	"todolist/model"
	"todolist/task"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
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
	token := jwt.New(jwt.GetSigningMethod(mw.SigningAlgorithm))
	claims := token.Claims.(jwt.MapClaims)

	if mw.PayloadFunc != nil {
		for key, value := range mw.PayloadFunc(data) {
			claims[key] = value
		}
	}

	expire := mw.TimeFunc().Add(mw.Timeout)
	claims["exp"] = expire.Unix()
	claims["orig_iat"] = mw.TimeFunc().Unix()
	tokenString, err := mw.signedString(token)
	if err != nil {
		mw.unauthorized(ctx, c, http.StatusUnauthorized, mw.HTTPStatusMessageFunc(ErrFailedTokenCreation, ctx, c))
		return
	}
}

func Logout(ctx context.Context, c *app.RequestContext) {

}
