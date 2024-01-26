package people

import (
	"people/internal/v1/service"
)

type PeopleHandler struct {
	PeopleService service.PeopleService
}

func NewPeopleHandler(personService service.PeopleService) *PeopleHandler {
	return &PeopleHandler{
		PeopleService: personService,
	}
}
