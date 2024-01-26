package app

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func SetupEnvironment() {
	err := godotenv.Load()
	if err != nil {
		logrus.Warn("Error loading .env file. Using default environment variables.")
	}
}
