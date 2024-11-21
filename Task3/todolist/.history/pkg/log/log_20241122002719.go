package log

import (
	"context"
	"time"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/logger/accesslog"
)

//默认日志格式
//[${time}] ${status} - ${latency} ${method} ${path}
//eg:
//[21:54:36] 200 - 2.906859ms GET /ping

func Mylog() app.HandlerFunc {
	return accesslog.New(
		//自定义时间格式
		accesslog.WithTimeFormat(time.RFC1123), //年月日时分秒
		//配置时间戳的刷新间隔
		accesslog.WithTimeInterval(time.Second),
		//自定义日志打印函数
		accesslog.WithAccessLogFunc(MyCtxInfof),
	)
}

func MyCtxInfof(ctx context.Context, format string, v ...interface{}) {
	logger.CtxInfof(ctx, format, v...)
}

type MyLoggger struct {
	hlog.FullLogger
}
