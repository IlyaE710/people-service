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

	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		logrus.Fatal(err)
	}

	hook := logrustash.New(conn, &logrus.JSONFormatter{})
	logrus.AddHook(hook)
}
