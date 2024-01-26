package people

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"people/internal/v1/handler/people/dto"
	serviceDro "people/internal/v1/service/people/dto"
)

// Update
// @Summary Обновление данных о человеке
// @Description Обновляет информацию о человеке на основе переданных данных.
// @Tags People
// @Accept json
// @Produce json
// @Param updateRequest body dto.UpdatePeopleRequest true "Данные для обновления пользователя"
// @Success 200
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /v1/people/ [put]
func (h *PeopleHandler) Update(c *gin.Context) {
	var updateRequest dto.UpdatePeopleRequest
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid request payload"})
		return
	}

	// Преобразование UpdatePeopleRequest в dto.UpdatePeople
	updateDto := serviceDro.UpdatePeople{
		ID:         updateRequest.ID,
		Name:       updateRequest.Name,
		Surname:    updateRequest.Surname,
		Age:        updateRequest.Age,
		Gender:     updateRequest.Gender,
		Patronymic: updateRequest.Patronymic,
		Country:    convertCountryRequests(updateRequest.Country),
	}

	logrus.Info("Преобразование UpdatePeopleRequest в dto.UpdatePeople", updateDto)
	// Вызов метода PeopleService.Update с dto.UpdatePeople
	err := h.PeopleService.Update(updateDto)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "Failed to update people"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "People updated successfully"})
}

// convertCountryRequests конвертирует []handlerDto.CountryRequest в []dto.Country.
func convertCountryRequests(countryRequests []dto.CountryRequest) []serviceDro.Country {
	var countries []serviceDro.Country

	for _, countryRequest := range countryRequests {
		country := serviceDro.Country{
			CountryID:   countryRequest.CountryID,
			Probability: countryRequest.Probability,
		}
		countries = append(countries, country)
	}

	return countries
}
