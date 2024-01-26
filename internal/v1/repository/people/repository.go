package people

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"people/internal/v1/repository/people/dto"
	"people/internal/v1/repository/people/entity"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetByName(name string) (*dto.People, error) {
	model := entity.People{}

	if err := r.db.Where("name = ?", name).Preload("Country").First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	result := &dto.People{
		ID:         model.ID,
		Name:       model.Name,
		Surname:    model.Surname,
		Patronymic: model.Patronymic,
		Age:        model.Age,
		Gender:     model.Gender,
		Country:    make([]dto.Country, 0),
	}

	for _, countryModel := range model.Country {
		result.Country = append(result.Country, dto.Country{
			ID:          countryModel.ID,
			CountryID:   countryModel.CountryID,
			Probability: countryModel.Probability,
		})
	}

	return result, nil
}

func (r *Repository) GetAll() ([]*dto.People, error) {
	var models []entity.People

	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}

	var result []*dto.People

	for _, model := range models {
		result = append(result, &dto.People{
			Name:   model.Name,
			Age:    model.Age,
			Gender: model.Gender,
		})
	}

	return result, nil
}

func (r *Repository) Delete(person dto.DeletePeople) error {
	return r.db.Where("id = ?", person.Id).Delete(&entity.People{}).Error
}

func (r *Repository) Update(person dto.UpdatePeople) error {
	model := entity.People{
		ID:         person.ID,
		Name:       person.Name,
		Surname:    person.Surname,
		Patronymic: person.Patronymic,
		Age:        person.Age,
		Gender:     person.Gender,
		Country:    convertCountry(person.Country),
	}

	logrus.Info("Преобразование repDto.UpdatePeople в entity.People", model)
	return r.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&model).Error
}

func (r *Repository) Add(person dto.CreatePeople) error {
	model := entity.People{
		Name:       person.Name,
		Surname:    person.Surname,
		Patronymic: person.Patronymic,
		Age:        person.Age,
		Gender:     person.Gender,
		Country:    convertCountry(person.Country),
	}
	return r.db.Create(&model).Error
}

func convertCountry(countryDto []dto.Country) []entity.Country {
	var countries []entity.Country

	for _, country := range countryDto {
		country := entity.Country{
			CountryID:   country.CountryID,
			Probability: country.Probability,
		}
		countries = append(countries, country)
	}

	return countries
}
