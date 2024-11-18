package Log

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	Out       io.Writer
	Hooks     LevelHooks
	Formatter Formatter
	// 其他字段...
}

func Init() {
	logger := logrus.New()

	out, err := os.OpenFile("file.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	logger.Out = out
}
