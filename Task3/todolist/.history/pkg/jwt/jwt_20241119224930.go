package jwt

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
)

func MyJwt() app.HandlerFunc {
	var newmiddleware = jwt.HertzJWTMiddleware{
		//设置密钥
		Key: []byte("mysecretkey"),
		//设置密钥函数
		KeyFunc: keyfunc,
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

func (t *jwt.Token) keyfunc(interface{}, error) {
	if jwt.GetSigningMethod(mw.SigningAlgorithm) != t.Method {
		return nil, ErrInvalidSigningAlgorithm
	}
	if mw.usingPublicKeyAlgo() {
		return mw.pubKey, nil
	}

	// save token string if valid
	c.Set("JWT_TOKEN", token)

	return mw.Key, nil
}

func (ctx context.Context, c *app.RequestContext) authenticator(interface{}, error) {
	var loginVals login
	if err := c.BindAndValidate(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	userID := loginVals.Username
	password := loginVals.Password

	if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
		return &User{
			UserName:  userID,
			LastName:  "Hertz",
			FirstName: "CloudWeGo",
		}, nil
	}

	return nil, jwt.ErrFailedAuthentication
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
