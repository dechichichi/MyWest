package userhandler

import (
	"context"
	"fmt"
	"net/http"
	JWT "todolist/middleware/jwt"
	"todolist/model"
	myutils "todolist/pkg"
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
		token, claims, err := myutils.GenerateToken(name)
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
	user, found := c.Get(JWT.IdentityKey)
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

func Login(ctx context.Context, c *app.RequestContext) {
	name := c.Query("username")
	pwd := c.Query("password")
	// 验证用户名密码
	err := task.TAsk(name, pwd)
	if err != nil {
		c.JSON(200, utils.H{
			"message": "用户名或密码错误",
			"code":    http.StatusBadRequest,
		})
		return
	}
	// 验证成功就登录成功
	// 接下来生成 token 返回给前端
	token, _, err := myutils.GenerateToken(name)
	if err != nil {
		c.JSON(200, utils.H{
			"message": "登录失败",
			"code":    http.StatusBadRequest,
		})
		return
	}
	claims, _ := myutils.GetClaimsFromJWT(ctx, c)
	c.JSON(http.StatusOK, utils.H{
		"message": "登录成功",
		"code":    http.StatusOK,
		"token":   token,
		"claims":  claims,
	})
}
func Logout(ctx context.Context, c *app.RequestContext) {
	// 验证token的有效性
	_, err := myutils.ParseToken(ctx, c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.H{
			"message": "无效的token",
			"code":    http.StatusUnauthorized,
		})
		return
	}
	// 清除token
	myutils.LogoutResponse(ctx, c, http.StatusOK)
	c.JSON(http.StatusOK, utils.H{
		"message": "退出成功",
		"code":    http.StatusOK,
	})
}
