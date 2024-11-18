package routes

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Router() {
	h := server.New(server.WithHostPorts(":8080"))
	//中间件
	h.Use()

	h.Spin()
}
