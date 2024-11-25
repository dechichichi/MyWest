package adminhandler

import (
	"context"
	"todolist/config"
	"todolist/model"

	"github.com/cloudwego/hertz/pkg/app"
	"gorm.io/gorm"
)

type Params struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var User model.User

func Login(c context.Context, ctx *app.RequestContext) {
	name := ctx.Query("username")
	pwd := ctx.Query("password")
	err := check(name, pwd)
	//验证用户名密码
}
func check(name, pwd string) error {
	db, err := gorm.Open("mysql", config.Dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return
}

func Logout(c context.Context, ctx *app.RequestContext) {

}
