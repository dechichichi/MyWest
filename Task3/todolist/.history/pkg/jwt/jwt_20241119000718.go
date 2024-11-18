package jwt

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
)

func MyJwt() []app.HandlerFunc {
	return []app.HandlerFunc{func(ctx context.Context, c *app.RequestContext) {
		fmt.Println("group middleware")
		c.Next(ctx)
	}}
}
