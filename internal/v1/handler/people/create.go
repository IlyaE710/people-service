package people

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"people/internal/v1/handler/people/dto"
	serviceDto "people/internal/v1/service/people/dto"
	"strings"
)

// Create
// @Summary Создание нового пользователя
// @Tags People
// @Description Создает нового пользователя с использованием предоставленных данных
// @Accept json
// @Produce json
// @Param request body dto.CreatePeopleRequest true "Данные для создания пользователя"
// @Success 201 "Success"
// @Failure 400 {object} dto.ErrorResponse "Bad Request"
// @Failure 500 {object} dto.ErrorResponse "Internal Server Error"
// @Router /v1/people [post]
func (h *PeopleHandler) Create(c *gin.Context) {
	var requestBody dto.CreatePeopleRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		logrus.Errorf("Error binding JSON request: %v", err)
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	logrus.WithFields(logrus.Fields{
		"Name":       requestBody.Name,
		"Surname":    requestBody.Surname,
		"Patronymic": requestBody.Patronymic,
	}).Info("Received CreatePeopleRequest")
	if err := h.PeopleService.Add(serviceDto.CreatePeople{
		Name:       strings.Title(requestBody.Name),
		Surname:    strings.Title(requestBody.Surname),
		Patronymic: strings.Title(requestBody.Patronymic),
	}); err != nil {
		logrus.Errorf("Error adding person: %v", err)
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	logrus.Info("Person added successfully")
	c.Status(http.StatusCreated)
}
