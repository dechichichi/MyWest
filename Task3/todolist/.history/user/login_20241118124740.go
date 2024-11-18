package user

import (
	"context"
	"net/http"
	"time"

	"todolist/model"
	"todolist/task"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
	"golang.org/x/crypto/bcrypt"
)

// 用户登录处理函数
func Login(w http.ResponseWriter, r *http.Request) {
	// 从请求中获取用户名和密码
	userName := r.FormValue("username")
	password := r.FormValue("password")

	// 验证用户凭据
	// 从数据库检索用户信息
	user, err := task.Ask(userName)
	if err != nil {
		// 处理错误
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		// 密码错误
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// 如果验证成功，创建 Token
	authMiddleware, err := jwt.New(&jwt.HertzJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour * 24, // 设置JWT的过期时间为24小时
		IdentityKey: "username",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					"username": v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			claims := jwt.ExtractClaims(ctx, c)
			id := claims["username"].(string)
			// 这里需要根据实际情况查询数据库获取用户信息
			// 假设我们已经有了用户ID，现在需要从数据库中获取用户对象
			user, err := task.AskByID(id)
			if err != nil {
				return nil, err
			}
			return user, nil
		},
	})
	if err != nil {
		// 处理错误
		http.Error(w, "Error creating token", http.StatusInternalServerError)
		return
	}

	// 发送 Token 给客户端
	tokenString, expire, err := authMiddleware.GenerateToken(user)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Authorization", tokenString)
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24), // 设置cookie过期时间为24小时
		HttpOnly: true,
		Secure:   true, // 如果是https站点，设置为true
	})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful"))
}
