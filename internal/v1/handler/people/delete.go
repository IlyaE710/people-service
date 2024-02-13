package people

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"people/internal/v1/handler/people/dto"
	serviceDto "people/internal/v1/service/people/dto"
	"strconv"
)

// Delete
// @Summary Удалить человека
// @Description Удаляет человека по указанному идентификатору.
// @ID delete-people
// @Produce json
// @Param id path int64 true "Идентификатор человека" Format(int64)
// @Success 204 {string} string "No Content"
// @Failure 400 {object} dto.ErrorResponse "Invalid ID format"
// @Failure 500 {object} dto.ErrorResponse "Internal Server Error"
// @Router /v1/people/{id} [delete]
// @Tags People
func (h *PeopleHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		logrus.Errorf("Invalid ID format: %v", err)
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid ID format"})
		return
	}
	logrus.WithFields(logrus.Fields{
		"ID": id,
	}).Info("Received DeletePeopleRequest")
	errS := h.PeopleService.Delete(serviceDto.DeletePeople{ID: id})
	if errors.Is(errS, gorm.ErrRecordNotFound) {
		logrus.Warnf("Record not found for deletion: ID %d", id)
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Message: "Record not found"})
		return
	}
	if errS != nil {
		logrus.Errorf("Error deleting person: %v", errS)
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "Internal Server Error"})
		return
	}

	logrus.Info("Person deleted successfully")
	c.Status(http.StatusNoContent)
}
