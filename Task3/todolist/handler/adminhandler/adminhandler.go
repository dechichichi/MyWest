package adminhandler

import (
	"context"
	"net/http"
	JWT "todolist/middleware/jwt"
	"todolist/task"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

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
	token, _, err := JWT.JwtMiddleware.TokenGenerator(name)
	if err != nil {
		c.JSON(200, utils.H{
			"message": "登录失败",
			"code":    http.StatusBadRequest,
		})
		return
	}
	claims, _ := JWT.JwtMiddleware.GetClaimsFromJWT(ctx, c)
	c.JSON(http.StatusOK, utils.H{
		"message": "登录成功",
		"code":    http.StatusOK,
		"token":   token,
		"claims":  claims,
	})
}
func Logout(ctx context.Context, c *app.RequestContext) {
	// 验证token的有效性
	_, err := JWT.JwtMiddleware.ParseToken(ctx, c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.H{
			"message": "无效的token",
			"code":    http.StatusUnauthorized,
		})
		return
	}
	// 清除token
	JWT.JwtMiddleware.LogoutResponse(ctx, c, http.StatusOK)
	c.JSON(http.StatusOK, utils.H{
		"message": "退出成功",
		"code":    http.StatusOK,
	})
}
