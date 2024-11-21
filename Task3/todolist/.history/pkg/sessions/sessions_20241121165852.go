package sessions

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/redis"
)

func MySession() app.HandlerFunc {
	//redis
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	return sessions.New("mysession", store)

}
