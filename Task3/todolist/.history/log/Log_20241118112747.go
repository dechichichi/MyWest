package main

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	Out       io.Writer
	Hooks     logrus.LevelHooks // 修改这里，使用 logrus 包中的 LevelHooks
	Formatter logrus.Formatter  // 这里也应确保使用 logrus 的 Formatter
	// 其他字段...
}

func Init() {
	logger := logrus.New()

	out, err := os.OpenFile("file.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		// 处理错误
	}
	logger.Out = out
	// 其他初始化代码...
}
