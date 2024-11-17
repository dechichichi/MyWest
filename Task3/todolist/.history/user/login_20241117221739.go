package user

import (
	"net/http"
	"time"
)

// 用户登录处理函数
func Login(w http.ResponseWriter, r *http.Request) {
	// 从请求中获取用户名和密码
	// ...

	// 验证用户凭据
	// ...

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
