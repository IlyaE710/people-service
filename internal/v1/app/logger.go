package app

import (
	"github.com/sirupsen/logrus"
	"os"
)

func SetupLogger() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logLevelStr := os.Getenv("LOG_LEVEL")
	logLevel, err := logrus.ParseLevel(logLevelStr)
	if err != nil {
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetLevel(logLevel)
	}
}
