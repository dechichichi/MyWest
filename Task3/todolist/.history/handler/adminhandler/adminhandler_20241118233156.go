package adminhandler

import "github.com/cloudwego/hertz/pkg/app"

type Params struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckParams(*app.RequestContext) error {

}

func Adminhandler(*app.HandlerFunc) {

}
