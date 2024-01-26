package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"people/internal/v1/handler/people"
)

func SetupServer(serviceLocator ServiceLocator) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(LoggingMiddleware())

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

	logrus.Info("Server setup complete")

	return r
}
