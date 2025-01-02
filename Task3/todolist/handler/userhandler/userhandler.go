package userhandler

import (
	"context"
	"fmt"
	"net/http"
	"todolist/middleware/jwt"
	"todolist/model"
	"todolist/task"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func Register(ctx context.Context, c *app.RequestContext) {
	name := c.Query("username")
	password := c.Query("password")
	email := c.Query("email")
	if _, err := task.Ask(name); err == nil {
		c.JSON(http.StatusBadRequest, utils.H{
			"message": "user already exists",
			"code":    http.StatusBadRequest,
		})
		return
	}
	if _, err := task.Add(name, password, email); err != nil {
		c.JSON(http.StatusInternalServerError, utils.H{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, utils.H{
		"message": "success",
		"code":    http.StatusOK,
	})
}

func Auth(ctx context.Context, c *app.RequestContext) {
	name := c.Query("username")
	password := c.Query("password")
	if user, err := task.Auth(name, password); err != nil {
		c.JSON(http.StatusUnauthorized, utils.H{
			"message": err.Error(),
			"code":    http.StatusUnauthorized,
		})
		return
	} else {
		token, claims, err := jwt.JwtMiddleware.TokenGenerator(name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.H{
				"message": "token generation failed",
				"code":    http.StatusInternalServerError,
			})
			return
		}
		c.JSON(http.StatusOK, utils.H{
			"message": "success",
			"code":    http.StatusOK,
			"data": utils.H{
				"username": user.UserName,
				"email":    user.Email,
				"token":    token,
				"claims":   claims,
			},
		})
	}
}

// 检查用户身份是否有效
func Ping(ctx context.Context, c *app.RequestContext) {
	user, found := c.Get(jwt.IdentityKey)
	if !found {
		c.JSON(http.StatusInternalServerError, utils.H{
			"message": "error retrieving user",
			"code":    http.StatusInternalServerError,
		})
		return
	}
	if user == nil {
		c.JSON(http.StatusUnauthorized, utils.H{
			"message": "unauthorized",
			"code":    http.StatusUnauthorized,
		})
		return
	}
	c.JSON(http.StatusOK, utils.H{
		"message": fmt.Sprintf("username:%v", user.(*model.User).UserName),
	})
}
