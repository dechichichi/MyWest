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

type HertzJWTMiddleware struct {
	Key     []byte
	Timeout time.Duration
}

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
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &model.User{
				ID: claims["username"].(string),
			}
		},
	})
	if err != nil {
		// 处理错误
		http.Error(w, "Error creating token", http.StatusInternalServerError)
		return
	}

	// 发送 Token 给客户端
	tokenString, err := authMiddleware.GenerateToken(user)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Authorization", tokenString)
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful"))
}

// GenerateToken 生成JWT Token
func (j *HertzJWTMiddleware) GenerateToken(identity interface{}) (string, error) {
	claims := jwt.MapClaims{
		"exp":      time.Now().Add(j.Timeout).Unix(),
		"iat":      time.Now().Unix(),
		"nbf":      time.Now().Unix(),
		"identity": identity,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.Key)
}
