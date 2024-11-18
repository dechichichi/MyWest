package adminhandler

import "github.com/cloudwego/hertz/pkg/app"

type Params struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckParams(params *Params) error {

}

func Adminhandler(*app.HandlerFunc) {

}
