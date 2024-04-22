package logger

import (
	"go-fiber-boilerplate/pkg/config"
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

var logger = &Logger{}

func GetLogger() *Logger {
	return logger
}

func SetupLogger() {
	logger = &Logger{logrus.New()}
	logger.Formatter = &logrus.JSONFormatter{}
	logger.SetOutput(os.Stdout)

	if config.AppConfig().Environtment == "development" {
		logger.SetLevel(logrus.DebugLevel)
	}
}
