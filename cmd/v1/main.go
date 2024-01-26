package main

import (
	"github.com/sirupsen/logrus"
	"os"
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

// @host      ${SWAGGER_HOST}
// @BasePath  ${SWAGGER_BASE_PATH}

func main() {
	app.SetupEnvironment()
	app.SetupLogger()

	db, err := app.SetupDatabase()
	if err != nil {
		logrus.Fatal("failed to set up database:", err)
	}

	serviceLocator := app.NewServiceLocator(db)

	r := app.SetupServer(*serviceLocator)

	logrus.Info("run server")

	addr := os.Getenv("SERVER_ADDRESS")
	if addr == "" {
		addr = ":8080"
	}

	logrus.Infof("Запуск сервера на адресе %s", addr)
	r.Run(addr)
}
