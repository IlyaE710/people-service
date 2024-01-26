package repository

import "people/pkg/repository/people/reponse"

type PeopleRepository interface {
	GetByName(name string) (*reponse.People, error)
}
