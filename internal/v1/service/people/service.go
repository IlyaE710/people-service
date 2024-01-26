package people

import (
	"github.com/sirupsen/logrus"
	"people/internal/v1/repository"
	"people/internal/v1/repository/people/dto"
	serviceDto "people/internal/v1/service/people/dto"
	externalRepository "people/pkg/repository"
)

type Service struct {
	peopleRepository   repository.PeopleRepository
	externalRepository externalRepository.PeopleRepository
}

func NewService(
	userRepository repository.PeopleRepository,
	externalRepository externalRepository.PeopleRepository,
) *Service {
	return &Service{
		peopleRepository:   userRepository,
		externalRepository: externalRepository,
	}
}

func (s Service) GetByName(name string) (*serviceDto.People, error) {
	person, err := s.peopleRepository.GetByName(name)

	if err != nil {
		return nil, err
	}

	result := &serviceDto.People{
		ID:         person.ID,
		Name:       person.Name,
		Surname:    person.Surname,
		Patronymic: person.Patronymic,
		Age:        person.Age,
		Gender:     person.Gender,
	}

	for _, countryModel := range person.Country {
		result.Country = append(result.Country, serviceDto.Country{
			ID:          countryModel.ID,
			CountryID:   countryModel.CountryID,
			Probability: countryModel.Probability,
		})
	}

	return result, nil
}

func (s Service) GetAll() ([]*serviceDto.People, error) {
	peopleList, err := s.peopleRepository.GetAll()

	if err != nil {
		return nil, err
	}

	var result []*serviceDto.People

	// Преобразование списка людей в список DTO
	for _, person := range peopleList {
		result = append(result, &serviceDto.People{
			Name:   person.Name,
			Age:    person.Age,
			Gender: person.Gender,
		})
	}

	return result, nil
}

func (s Service) Delete(people serviceDto.DeletePeople) error {

	return s.peopleRepository.Delete(dto.DeletePeople{
		Id: people.ID,
	})
}

func (s Service) Update(people serviceDto.UpdatePeople) error {
	p := dto.UpdatePeople{
		ID:         people.ID,
		Name:       people.Name,
		Surname:    people.Name,
		Patronymic: people.Name,
		Age:        people.Age,
		Gender:     people.Gender,
		Country:    convertCountry(people.Country),
	}
	logrus.Info("Преобразование serviceDto.UpdatePeople в repDto.UpdatePeople", p)
	return s.peopleRepository.Update(p)
}

func (s Service) Add(people serviceDto.CreatePeople) error {
	response, err := s.externalRepository.GetByName(people.Name)
	if err != nil {
		return err
	}

	countryInfoList := make([]dto.Country, 0, len(response.Country))
	for _, countryInfo := range response.Country {
		countryInfoList = append(countryInfoList, dto.Country{
			CountryID:   countryInfo.CountryID,
			Probability: countryInfo.Probability,
		})
	}

	return s.peopleRepository.Add(dto.CreatePeople{
		Name:       response.Name,
		Surname:    people.Surname,
		Patronymic: people.Patronymic,
		Gender:     response.Gender,
		Age:        response.Age,
		Country:    countryInfoList,
	})
}

func convertCountry(countryDto []serviceDto.Country) []dto.Country {
	var countries []dto.Country

	for _, country := range countryDto {
		country := dto.Country{
			CountryID:   country.CountryID,
			Probability: country.Probability,
		}
		countries = append(countries, country)
	}

	return countries
}
