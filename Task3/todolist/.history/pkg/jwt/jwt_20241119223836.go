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
func authorizator(c context.Context) (interface{}, error) {
	return nil, nil
}
func authorizator(c context.Context) (interface{}, error) {
	return nil, nil
}
