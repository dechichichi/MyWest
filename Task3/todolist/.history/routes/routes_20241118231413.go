package routes

import (
	_ "ToDoList/docs"
	"ToDoList/handler/adminHandler"
	"ToDoList/handler/taskHandler"
	"ToDoList/handler/userHandler"
	jwt "ToDoList/pkg/authorization"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/logger/accesslog"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/cookie"
	"github.com/hertz-contrib/swagger"
	swaggerFiles "github.com/swaggo/files"
)

func Router() *server.Hertz {
	//session扩展
	h := server.New(server.WithHostPorts(":8000"))
	h.Use(accesslog.New())
	//store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("PANDORA PARADOXXX"))
	store := cookie.NewStore([]byte("PANDORA PARADOXXX"))
	h.Use(sessions.New("sessionId", store))

	mw := jwt.JWT()

	v1 := h.Group("todolist/")
	{
		v1.POST("user/register", userHandler.UserRegister)
		v1.POST("user/login", mw.LoginHandler) //Hertz JWT:Authenticator
	}
	auth := v1.Group("auth")
	auth.Use(mw.MiddlewareFunc())
	{
		//auth.GET("/test", jwt.TestHandler)
		//实现id访问
		auth.GET("/:id", userHandler.IdLogin) //test
		auth.POST("/:id/task/create", taskHandler.CreateTask)
		auth.GET("/:id/task/:tid", taskHandler.CheckTask)
		auth.POST("/:id/tasks", taskHandler.ListTask)
		auth.POST("/:id/task/search", taskHandler.SearchTask)
		auth.PUT("/:id/task/:tid", taskHandler.UpdateTask)
		auth.DELETE("/:id/task/:tid", taskHandler.DeleteTask)

	}
	admin := v1.Group("admin")
	admin.Use(mw.MiddlewareFunc())
	{
		admin.POST("/listusers", adminHandler.ListUsers)
		admin.GET("/add/:id", adminHandler.AddAdmin)
		admin.GET("/block/:id", adminHandler.BlockUser)
	}
	url := swagger.URL("http://localhost:8000/swagger/doc.json") // The url pointing to API definition
	h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler, url))
	return h
}
