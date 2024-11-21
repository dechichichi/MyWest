package keyauth

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/keyauth"
)

func Init() {
	h := server.Default()
	h.Use(keyauth.New(
		keyauth.WithValidator(func(ctx context.Context, requestContext *app.RequestContext, s string) (bool, error) {
			if s == "test_admin" {
				return true, nil
			}
			return false, nil
		}),
	))
	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		value, _ := c.Get("token")
		c.JSON(consts.StatusOK, utils.H{"ping": value})
	})
	h.Spin()
}
