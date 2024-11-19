package jwt

import (
	"context"

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
		//设置授权已认证的用户路由访问权限
		Authorizator: authorizator,
		//设置登陆成功后为向 token 中添加自定义负载信息
		PayloadFunc: payloadFunc,
		//设置 jwt 验证流程失败处理函数
		Unauthorized: unauthorized,
		//设置登录
		LoginResponse: loginResponse,
		//设置登出
		LogoutResponse: logoutResponse,
		//检索用户信息
		IdentityKey: "id",
		//获取身份信息
		IdentityHandler: identityHandler,
		//设置cookie名称
		CookieName: "jwt",
	}
}

// KeyFunc 只在解析 token 时生效，签发 token 时不生效
