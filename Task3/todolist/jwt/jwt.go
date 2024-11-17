package jwtMaker

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Payload 自定义的JWT载荷
type Payload struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// JWTMaker 结构体，包含密钥和签名方法
type JWTMaker struct {
	signingKey []byte
}

// NewJWTMaker 创建一个新的JWTMaker实例
func NewJWTMaker(signingKey string) *JWTMaker {
	return &JWTMaker{
		signingKey: []byte(signingKey),
	}
}

// CreateToken 创建一个新的JWT
func (maker *JWTMaker) CreateToken(username string, expiration time.Duration) (string, error) {
	payload := &Payload{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString(maker.signingKey)
}

// VerifyToken 验证一个JWT
func (maker *JWTMaker) VerifyToken(tokenString string) (*Payload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Payload{}, func(token *jwt.Token) (interface{}, error) {
		return maker.signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Payload)
	if !ok || !token.Valid {
		return nil, jwt.NewValidationError("token invalid", jwt.ValidationErrorSignatureInvalid)
	}

	return claims, nil
}
