package jwt

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt/v4"
	"github.com/hertz-contrib/jwt"
)

func MyJwt() app.HandlerFunc {
	var newmiddleware = jwt.HertzJWTMiddleware{
		//设置密钥
		Key: []byte("mysecretkey"),
		//设置密钥函数
		KeyFunc: func(token *jwt.Token) (interface{}, error) {
			return []byte("mysecretkey"), nil
		},
		//设置登录时认证用户信息
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			return nil, nil
		},
		//设置登陆成功后为向 token 中添加自定义负载信息
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			return jwt.MapClaims{}
		},
		//获取身份信息
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			return nil
		},
		//设置cookie名称
		CookieName: "jwt",
		//设置登录的响应函数
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {

		},
		//设置登出的响应函数
		LogoutResponse: func(ctx context.Context, c *app.RequestContext, code int) {

		},
	}
}

// KeyFunc 只在解析 token 时生效，签发 token 时不生效
