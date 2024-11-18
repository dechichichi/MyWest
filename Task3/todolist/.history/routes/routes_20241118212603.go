package routes

import "github.com/cloudwego/hertz/pkg/app/server"
import ""github.com/hertz-contrib/logger/accesslog""
func Router() {
	h := server.New(server.WithHostPorts(":8080"))
	h.Use(accesslog.)
}