package userhandler

import (
	"context"
	"fmt"
	"todolist/model"
	"todolist/pkg/jwt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func Ping(ctx context.Context, c *app.RequestContext) {
	user, _ := c.Get(jwt.IdentityKey)
	c.JSON(200, utils.H{
		"message": fmt.Sprintf("username:%v", user.(*model.User).UserName),
	})
}
