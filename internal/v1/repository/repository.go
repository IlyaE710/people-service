package repository

import "people/internal/v1/repository/people/dto"

type PeopleRepository interface {
	GetByName(name string) (*dto.People, error)
	Delete(people dto.DeletePeople) error
	Update(people dto.UpdatePeople) error
	Add(people dto.CreatePeople) error
	GetAll() ([]*dto.People, error)
	GetById(id int) (*dto.People, error)
	DeleteCountry(people dto.DeletePeople) error
}
