package user

import (
	"net/http"
	"time"
	jwtMaker "todolist/jwt"
	"todolist/service"
	"todolist/task"

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
	// 设置JWT的过期时间为24小时
	expirationTime := time.Hour * 24
	// 创建一个新的JWTMaker实例
	tokener := jwtMaker.NewJWTMaker("news")
	// 创建一个新的JWT，设置用户名和过期时间
	token, err := tokener.CreateToken(userName, expirationTime)
	if err != nil {
		// 处理错误
		http.Error(w, "Error creating token", http.StatusInternalServerError)
		return
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
	token := r.Header.Get("Authorization")
	// 创建一个新的JWTMaker实例
	tokener := jwtMaker.NewJWTMaker("news")
	claims, err := tokener.VerifyToken(token)
	if err != nil {
		// Token 无效，拒绝访问
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	} else {
		// Token 有效，处理请求
		service.HandleRequest(w, r, claims)
	}
}
