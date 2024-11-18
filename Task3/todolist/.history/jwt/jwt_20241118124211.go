package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
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
