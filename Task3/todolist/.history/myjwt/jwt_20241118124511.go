package myjwt

import (
	"context"
	"net/http"
	"time"

	"github.com/hertz-contrib/jwt"
)

// GenerateToken 生成JWT Token
func (j *HertzJWTMiddleware) GenerateToken(identity interface{}) (string, error) {
	claims := jwt.MapClaims{
		"exp":      time.Now().Add(j.Timeout).Unix(),
		"iat":      time.Now().Unix(),
		"nbf":      time.Now().Unix(),
		"identity": identity,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.Key)
}

// Decode 解析并验证JWT Token
func (j *HertzJWTMiddleware) Decode(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return j.Key, nil
	})
	if err != nil {
		return nil, nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, nil, jwt.NewValidationError("token is invalid", jwt.ValidationErrorSignatureInvalid)
	}
	return token, claims, nil
}

// IdentityHandler 从请求中提取身份信息
func (j *HertzJWTMiddleware) IdentityHandler(ctx context.Context) (interface{}, error) {
	req := ctx.Value("request").(*http.Request)
	tokenString := req.Header.Get("Authorization")
	if tokenString == "" {
		return nil, jwt.NewValidationError("Authorization header is missing", jwt.ValidationErrorExpired)
	}
	// 解析Token
	token, claims, err := j.Decode(tokenString)
	if err != nil {
		return nil, err
	}
	identity, ok := claims["identity"]
	if !ok {
		return nil, jwt.NewValidationError("identity is missing", jwt.ValidationErrorExpired)
	}
	return identity, nil
}
