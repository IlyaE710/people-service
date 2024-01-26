package app

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	personRepository "people/internal/v1/repository/people"
	"people/internal/v1/repository/people/entity"
	"people/internal/v1/service"
	personService "people/internal/v1/service/people"
	externalRepository "people/pkg/repository/people"
)

type serviceLocator struct {
	PeopleService service.PeopleService
	Db            *gorm.DB
}

func NewServiceLocator() *serviceLocator {
	dsn := "host=localhost user=user password=pass dbname=db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&entity.People{}, &entity.Country{})

	r := personRepository.NewRepository(db)
	eR := externalRepository.NewRepository()
	return &serviceLocator{
		PeopleService: personService.NewService(r, eR),
		Db:            db,
	}
}
