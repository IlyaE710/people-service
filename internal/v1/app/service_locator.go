package app

import (
	"gorm.io/gorm"
	personRepository "people/internal/v1/repository/people"
	"people/internal/v1/service"
	personService "people/internal/v1/service/people"
	externalRepository "people/pkg/repository/people"
)

type ServiceLocator struct {
	PeopleService service.PeopleService
	db            *gorm.DB
}

func NewServiceLocator(db *gorm.DB) *ServiceLocator {
	r := personRepository.NewRepository(db)
	eR := externalRepository.NewRepository()
	return &ServiceLocator{
		PeopleService: personService.NewService(r, eR),
		db:            db,
	}
}
