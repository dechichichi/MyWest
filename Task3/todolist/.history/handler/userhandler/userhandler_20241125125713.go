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
		panic(err)
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
