package sessions

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/redis"
)

func Init() {
	h := server.Default(server.WithHostPorts(":8000"))
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	h.Use(sessions.New("mysession", store))

	h.GET("/incr", func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, utils.H{"count": count})
	})
	h.Spin()
}
