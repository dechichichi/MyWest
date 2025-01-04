package myutils

import (
	"context"
	"time"
	JWT "todolist/middleware/jwt"

	"github.com/cloudwego/hertz/pkg/app"
	JWTv4 "github.com/golang-jwt/jwt/v4"
	"github.com/hertz-contrib/jwt"
)

func GenerateToken(username string) (string, time.Time, error) {
	return JWT.JwtMiddleware.TokenGenerator(username)
}

func ParseToken(ctx context.Context, c *app.RequestContext) (*JWTv4.Token, error) {
	return JWT.JwtMiddleware.ParseToken(ctx, c)
}
func GetClaimsFromJWT(ctx context.Context, c *app.RequestContext) (jwt.MapClaims, error) {
	return JWT.JwtMiddleware.GetClaimsFromJWT(ctx, c)
}

func LogoutResponse(ctx context.Context, c *app.RequestContext, code int) {
	JWT.JwtMiddleware.LogoutResponse(ctx, c, code)
}
