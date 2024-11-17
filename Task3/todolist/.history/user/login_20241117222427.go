package user

import (
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// 用户登录处理函数
func Login(w http.ResponseWriter, r *http.Request) {
	// 从请求中获取用户名和密码
	userName := r.FormValue("username")
	password := r.FormValue("password")

	// 验证用户凭据
	// 从数据库检索用户信息
	user, err := getUserByUsername(userName)
	if err != nil {
		// 处理错误
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		// 密码不匹配
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// 如果验证成功，创建 Token
	token, err := jwtMaker.CreateToken(userName, time.Hour)
	if err != nil {
		// 处理错误
		// ...
	} else {
		// 发送 Token 给客户端
		http.SetCookie(w, &http.Cookie{
			Name:    "auth_token",
			Value:   token,
			Expires: time.Now().Add(time.Hour),
			// ...
		})
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Login successful"))
	}
}

// 需要认证的路由处理函数
func SecureRoute(w http.ResponseWriter, r *http.Request) {
	// 从请求中获取 Token
	// ...

	// 验证 Token
	claims, err := jwtMaker.VerifyToken(token)
	if err != nil {
		// Token 无效，返回错误
		// ...
	} else {
		// Token 有效，处理请求
		// ...
	}
}

func getUserByUsername(username string) (*User, error) {
	// 从数据库中检索用户信息
}
