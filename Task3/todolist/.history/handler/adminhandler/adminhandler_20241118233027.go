package adminhandler

import "github.com/cloudwego/hertz/pkg/app"

type Params struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func checkParams(params *Params) error {

}

func Adminhandler(*app.HandlerFunc) {

}
