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
		logrus.Errorf("Invalid request payload: %v", err)
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid request payload"})
		return
	}

	logrus.WithFields(logrus.Fields{
		"ID":         updateRequest.ID,
		"Name":       updateRequest.Name,
		"Surname":    updateRequest.Surname,
		"Patronymic": updateRequest.Patronymic,
	}).Info("Received UpdatePeopleRequest")

	updateDto := serviceDro.UpdatePeople{
		ID:         updateRequest.ID,
		Name:       updateRequest.Name,
		Surname:    updateRequest.Surname,
		Patronymic: updateRequest.Patronymic,
	}

	logrus.WithFields(logrus.Fields{
		"UpdateDto": updateDto,
	}).Info("Converted UpdatePeopleRequest to dto.UpdatePeople")

	err := h.PeopleService.Update(updateDto)
	if err != nil {
		logrus.Errorf("Failed to update people: %v", err)
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "Failed to update people"})
		return
	}

	logrus.Info("People updated successfully")

	c.JSON(http.StatusOK, gin.H{"message": "People updated successfully"})
}
