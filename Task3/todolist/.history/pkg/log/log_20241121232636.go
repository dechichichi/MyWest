package log

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/logger/accesslog"
)

func Mylog() app.HandlerFunc {
	return accesslog.New()
}
