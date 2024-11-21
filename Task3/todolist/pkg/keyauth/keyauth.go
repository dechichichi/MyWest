package keyauth

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/keyauth"
)

func MyKeyauth() app.HandlerFunc {
	return keyauth.New(
		//设置自定义的校验逻辑用于 token 校验
		keyauth.WithValidator(func(ctx context.Context, requestContext *app.RequestContext, s string) (bool, error) {
			if s == "test_admin" {
				return true, nil
			}
			return false, nil
		}),
		//设置校验 token 通过的自定义处理逻辑
		keyauth.WithSuccessHandler(func(ctx context.Context, c *app.RequestContext) {
			c.Next(ctx)
		}),
		//设置存储在请求上下文的 token 对应的 key
		keyauth.WithContextKey("token"),
	)
}
