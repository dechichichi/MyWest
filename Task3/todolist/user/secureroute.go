package user

import (
	"context"
	"net/http"
	"todolist/service"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
)

// 需要认证的路由处理函数
func SecureRoute(w http.ResponseWriter, r *http.Request) {
	// 从请求中获取 Token
	token := r.Header.Get("Authorization")
	// 创建一个新的JWTMaker实例
	authMiddleware, err := jwt.New(&jwt.HertzJWTMiddleware{
		Realm: "test zone",
		Key:   []byte("secret key"),
	})
	if err != nil {
		http.Error(w, "Invalid token", http.StatusInternalServerError)
		return
	}

	claims, err := authMiddleware.Authenticator(token, &app.RequestContext{Request: r})
	if err != nil {
		// Token 无效，拒绝访问
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Token 有效，处理请求
	user, err := authMiddleware.IdentityHandler(context.Background(), &app.RequestContext{Request: r})
	if err != nil {
		http.Error(w, "Error retrieving user identity", http.StatusInternalServerError)
		return
	}
	service.HandleRequest(w, r, user)
}
