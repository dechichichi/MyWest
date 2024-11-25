package userhandler

import (
	"context"
	"net/http"
	"todolist/task"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func Register(ctx context.Context, c *app.RequestContext) {
	name := c.Query("username")
	passward := c.Query("passward")
	email := c.Query("email")
	_, err := task.Ask(name)
	if err != nil {
		c.JSON(http.StatusOK, utils.H{
			"message": "user already exists",
			"code":    http.StatusBadRequest,
		})
		return
	}
	_, err = task.Add(name, passward, email)
	if err != nil {
		c.JSON(http.StatusOK, utils.H{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
		})
	}
	c.JSON(http.StatusOK, utils.H{
		"message": "success",
		"code":    http.StatusOK,
	})
}

func Auth(ctx context.Context, c *app.RequestContext) {
	name := c.Query("username")
	passward := c.Query("passward")
	user, err := task.Auth(name, passward)
	if err != nil {
		c.JSON(http.StatusOK, utils.H{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
		})
		return
	}
	c.JSON(http.StatusOK, utils.H{
		"message": "success",
		"code":    http.StatusOK,
		"data": utils.H{
			"username": user.Name,
			"email":    user.Email,
		},
	})
}
