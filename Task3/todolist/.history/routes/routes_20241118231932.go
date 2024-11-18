package routes

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Router() {
	h := server.New(server.WithHostPorts(":8080"))
	h1 := h.Group("/admin")
	h2 := h.Group("/item")
	//中间件
	h.Use()
	//

	h.Spin()
}

//http部分由三部分框架组成
