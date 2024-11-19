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

func keyfunc(c context.Context) (interface{}, error) {
	return nil, nil
}

func authenticator(c context.Context) (interface{}, error) {
	return nil, nil
}

func authorizator(c context.Context) (interface{}, error) {
	return nil, nil
}

func payloadFunc(c context.Context) (interface{}, error) {
	return nil, nil
}
func unauthorized(c context.Context) (interface{}, error) {
	return nil, nil
}
func loginResponse(c context.Context) (interface{}, error) {
	return nil, nil
}
func logoutResponse(c context.Context) (interface{}, error) {
	return nil, nil
}

func identityHandler(c context.Context) (interface{}, error) {
	return nil, nil
}
