package main

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	Out       io.Writer
	Hooks     logrus.LevelHooks
	Formatter logrus.Formatter
}

func Init() {
	logger := logrus.New()

	out, err := os.OpenFile("file.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	logger.Out = out
}
