package user

import (
	"context"
	MyJwt "todolist/jwt"
	"todolist/model"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
)

//JWT: A.B.C
//A: Header  固定
//B: Payload  存放信息
//C: Signature  签名 A与B加上密钥
//jwt就是验证C部分是否合法

func Login(username, password string) (string, error) {
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
			return &User{
				UserName: claims[identityKey].(string),
			}
		}
	})

}