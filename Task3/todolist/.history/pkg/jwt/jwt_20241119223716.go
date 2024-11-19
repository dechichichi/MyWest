package jwt

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
)

func MyJwt() app.HandlerFunc {
	var newmiddleware = jwt.HertzJWTMiddleware{
		//设置密钥
		key: "mysecretkey",
		//设置密钥函数
		keyFunc: keyfunc,
		//设置登录时认证用户信息
		Authenticator: authenticator,
	}
}

func keyfunc(c context.Context) (interface{}, error) {
	// 你可以在这里实现身份验证逻辑
	return nil, nil
}

func authenticator(c context.Context) (interface{}, error) {
	// 你可以在这里实现身份验证逻辑
	return nil, nil
}
