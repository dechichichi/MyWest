package userhandler

import "github.com/cloudwego/hertz/pkg/app"

type User{
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email"`
	Password string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}



func Register(c context.Context, ctx *app.RequestContext) {

}