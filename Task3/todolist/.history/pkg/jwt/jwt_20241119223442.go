package jwt

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
)

func MyJwt() app.HandlerFunc {
	var newmiddleware = jwt.HertzJWTMiddleware{
		//所属领域名称默认
		//签名算法默认
		key: "mysecretkey",
		keyFunc: func(token *jwt.Token) (interface{}, error) {
			return []byte("mysecretkey"), nil
		},
		//token 过期时间默认
		//最大token刷新时间默认
		Authenticator: authenticator,
	}
}

func (token *jwt.Token) authenticator(interface{}, error) {

}
