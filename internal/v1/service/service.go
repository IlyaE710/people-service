package service

import (
	"people/internal/v1/service/people/dto"
)

type PeopleService interface {
	GetByName(name string) (*dto.People, error)
	Delete(people dto.DeletePeople) error
	Update(people dto.UpdatePeople) error
	Add(people dto.CreatePeople) error
	GetAll() ([]*dto.People, error)
}
