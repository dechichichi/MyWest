package user

import (
	"context"
	"time"
	MyJwt "todolist/jwt"
	"todolist/model"
	"todolist/task"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/hertz-contrib/jwt"
)

// JWT: A.B.C
// A: Header  固定
// B: Payload  存放信息
// C: Signature  签名 A与B加上密钥
// jwt就是验证C部分是否合法
const identityKey = "username"

func Login(username, password string) (string, error) {
	// 验证用户名密码
	user, err := task.Ask(username)
	if err != nil {
		panic(err)
	}
	authMiddleware, err := jwt.New(&jwt.HertzJWTMiddleware{
		Key: []byte(MyJwt.SecretKey),
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			if (username == "admin" && password == "admin") || (username == "test" && password == "test") {
				return &model.User{
					UserName: username,
					Password: password,
				}, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			if v, ok := data.(*model.User); ok && v.UserName == "admin" {
				return true
			}

			return false
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			return jwt.MapClaims{}
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &model.User{
				UserName: claims[identityKey].(string),
			}
		},
		SendCookie:     true,
		TokenLookup:    "cookie: jwt",
		CookieMaxAge:   time.Hour,
		SecureCookie:   false,
		CookieHTTPOnly: false,
		CookieDomain:   ".test.com",
		CookieName:     "jwt-cookie",
		CookieSameSite: protocol.CookieSameSiteDisabled,
	})
	if err != nil {
		return "", err
	}
	authMiddleware.Timeout = time.Minute * 5
}
