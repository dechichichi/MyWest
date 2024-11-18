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
	var ctx context.Context
	ctx = context.WithValue(ctx, "auth", authMiddleware)
	claims, err := authMiddleware.Authenticator(ctx, &app.RequestContext{Request: r})
	if err != nil {
		// Token 无效，拒绝访问
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Token 有效，处理请求
	claims, err := authMiddleware.Authenticator(ctx, &app.RequestContext{Request: r})
	if err != nil {
		// Token 无效，拒绝访问
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// 尝试将 claims 类型断言为 map[string]interface{}
	claimsMap, ok := claims.(map[string]interface{})
	if !ok {
		// 如果断言失败，返回错误
		http.Error(w, "Invalid claims format", http.StatusUnauthorized)
		return
	}

	// 然后可以安全地索引
	user := claimsMap["user"].(string)
	service.HandleRequest(w, r, user)

}
