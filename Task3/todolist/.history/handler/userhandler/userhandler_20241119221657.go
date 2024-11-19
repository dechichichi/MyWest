package userhandler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

type User struct {
	Name            string `json:"name"`
	Age             int    `json:"age"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

func Register(c context.Context, ctx *app.RequestContext) {
	name := ctx.Query("name")
	age := ctx.Query("age")
	email := ctx.Query("email")
	password := ctx.Query("password")
	passwordConfirm := ctx.Query("password_confirm")
}
