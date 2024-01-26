package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "people/docs"
	"people/internal/v1/app"
	"people/internal/v1/handler/people"
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
	r := gin.Default()
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
	serviceLocator := app.NewServiceLocator()
	db, err := serviceLocator.Db.DB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	peopleHandler := people.NewPeopleHandler(serviceLocator.PeopleService)

	v1 := r.Group("api/v1")
	{
		v1people := v1.Group("/people")
		{
			v1people.POST("/", peopleHandler.Create)
			v1people.GET("/", peopleHandler.GetAll)
			v1people.GET("/:name", peopleHandler.GetPeopleByName)
			v1people.PUT("/", peopleHandler.Update)
			v1people.DELETE("/:id", peopleHandler.Delete)
		}
	}

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	logrus.Info("Запуск сервера")
	r.Run(":8080")
}
