package main

import (
	"github.com/sirupsen/logrus"
	_ "people/docs"
	"people/internal/v1/app"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api

func main() {
	app.SetupLogger()

	db, err := app.SetupDatabase()
	if err != nil {
		logrus.Fatal("failed to set up database:", err)
	}

	serviceLocator := app.NewServiceLocator(db)

	r := app.SetupServer(*serviceLocator)

	logrus.Info("run server")
	r.Run(":8080")
}
