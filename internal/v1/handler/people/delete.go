package people

import (
	"errors"
	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid ID format"})
		return
	}
	errS := h.PeopleService.Delete(serviceDto.DeletePeople{ID: id})
	if errors.Is(errS, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Message: "Record not found"})
		return
	}
	if errS != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "Internal Server Error"})
		return
	}

	c.Status(http.StatusNoContent)
}
