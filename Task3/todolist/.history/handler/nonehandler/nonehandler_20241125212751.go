package nonehandler

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func BlockUser(ctx context.Context, c *app.RequestContext) {
	if !Login(ctx, c) {
		c.JSON(200, utils.H{
			"message": "登录失败",
			"code":    http.StatusBadRequest,
		})
		return
	}
	id := c.Param("id")
	var blockUser AdminService.BlockUserService
	res := blockUser.Block(id)
	c.JSON(200, res)
}

func ListUsers(ctx context.Context, c *app.RequestContext) {
	if Login(ctx, c) != nil {
		c.JSON(200, utils.H{
			"message": "登录失败",
			"code":    http.StatusBadRequest,
		})
		return
	}
	var listUsers AdminService.ListUsersService
	if err := c.BindAndValidate(&listUsers); err != nil {
		c.JSON(e.InvalidParams, model.ErrorResponse{
			Status: e.ERROR,
		})
	} else {
		res := listUsers.List()
		c.JSON(e.SUCCESS, res)
	}
}
func AddAdmin(ctx context.Context, c *app.RequestContext) {
	if Login(ctx, c) != nil {
		c.JSON(200, utils.H{
			"message": "登录失败",
			"code":    http.StatusBadRequest,
		})
		return
	}
	id := c.Param("id")
	var addAdmin AdminService.AddAdminService
	res := addAdmin.AddAdmin(id)
	c.JSON(200, res)
}
