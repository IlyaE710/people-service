package app

import (
	"github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/sirupsen/logrus"
	"net"
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

	host, port := os.Getenv("LOG_HOST"), os.Getenv("LOG_PORT")
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		logrus.Error(err)
		return
	}

	hook := logrustash.New(conn, &logrus.JSONFormatter{})
	logrus.AddHook(hook)
}
