package people

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"people/internal/v1/handler/people/dto"
	_ "people/internal/v1/handler/people/dto"
	"strings"
)

// GetPeopleByName godoc
// @Summary Получение информации о человеке по имени
// @Description Получает информацию о человеке по его имени.
// @Tags People
// @ID get-people-by-name
// @Produce json
// @Param name path string false "Имя человека"
// @Success 200 {object} dto.PeopleResponse "Успешный ответ"
// @Failure 400 {object} dto.ErrorResponse "Ошибка запроса"
// @Failure 404 {object} dto.ErrorResponse "Человек не найден"
// @Failure 500 {object} dto.ErrorResponse "Внутренняя ошибка сервера"
// @Router /v1/people/{name} [get]
func (h *PeopleHandler) GetPeopleByName(c *gin.Context) {
	name := strings.Title(c.Param("name"))

	logrus.WithFields(logrus.Fields{
		"Name": name,
	}).Info("Received GetPeopleByNameRequest")

	person, err := h.PeopleService.GetByName(name)
	if err != nil {
		logrus.Errorf("Error retrieving person by name: %v", err)
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "Internal Server Error"})
		return
	}

	if person == nil {
		logrus.Warnf("Person not found by name: %s", name)
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Message: "Person not found"})
		return
	}

	result := &dto.PeopleResponse{
		Name:       person.Name,
		Surname:    person.Surname,
		Patronymic: person.Patronymic,
		Age:        person.Age,
		Gender:     person.Gender,
	}

	for _, countryModel := range person.Country {
		result.Country = append(result.Country, dto.CountryResponse{
			ID:          countryModel.ID,
			CountryID:   countryModel.CountryID,
			Probability: countryModel.Probability,
		})
	}

	logrus.WithFields(logrus.Fields{
		"Name":       person.Name,
		"Surname":    person.Surname,
		"Patronymic": person.Patronymic,
		"Age":        person.Age,
		"Gender":     person.Gender,
	}).Info("Successful response in GetPeopleByName handler")

	logrus.Info("Person get successfully")
	c.JSON(http.StatusOK, result)
}

// GetAll
// @Summary Получение списка всех людей
// @Description Получает список всех людей.
// @Tags People
// @ID get-all-people
// @Produce json
// @Success 200 {array} []dto.PeopleResponse "Успешный ответ"
// @Failure 404 {object} dto.ErrorResponse "Человек не найден"
// @Failure 500 {object} dto.ErrorResponse "Внутренняя ошибка сервера"
// @Router /v1/people/ [get]
func (h *PeopleHandler) GetAll(c *gin.Context) {
	logrus.Info("Received GetAllPeopleRequest")
	persons, err := h.PeopleService.GetAll()
	if err != nil {
		logrus.Errorf("Error retrieving all people: %v", err)
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "Internal Server Error"})
		return
	}

	if persons == nil {
		logrus.Warn("No people found")
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Message: "Person not found"})
		return
	}

	logrus.Info("Persons get successfully")
	c.JSON(http.StatusOK, persons)
}
