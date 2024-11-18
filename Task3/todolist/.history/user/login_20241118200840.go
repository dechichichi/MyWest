package user

import (
	"github.com/hertz-contrib/jwt"
)

//JWT: A.B.C
//A: Header  固定
//B: Payload  存放信息
//C: Signature  签名 A与B加上密钥
//jwt就是验证C部分是否合法

func Login(username, password string) (string, error) {
	authMiddleware, err := jwt.New(&jwt.HertzJWTMiddleware{})

}
