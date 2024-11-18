package handler

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz-examples/bizdemo/hertz_jwt/biz/model"
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_jwt/biz/mw"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

// Ping .
func Ping(ctx context.Context, c *app.RequestContext) {
	user, _ := c.Get(mw.IdentityKey)
	c.JSON(200, utils.H{
		"message": fmt.Sprintf("username:%v", user.(*model.User).UserName),
	})
}
