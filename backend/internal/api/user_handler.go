package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"modern-stack/internal/errors"
	"modern-stack/internal/models"
	"modern-stack/internal/services"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) handleError(c *gin.Context, err error) {
	switch err {
	case errors.ErrUserNotFound:
		c.JSON(http.StatusNotFound, models.NewResponse(
			http.StatusNotFound,
			err.Error(),
			nil,
		))
	default:
		c.JSON(http.StatusInternalServerError, models.NewResponse(
			http.StatusInternalServerError,
			err.Error(),
			nil,
		))
	}
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID := c.GetUint("user_id")
	user, err := h.userService.GetProfile(userID)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, models.NewResponse(
		http.StatusOK,
		"Profile retrieved successfully",
		user,
	))
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.NewResponse(
			http.StatusBadRequest,
			err.Error(),
			nil,
		))
		return
	}

	userID := c.GetUint("user_id")
	if err := h.userService.UpdateProfile(userID, req.Name); err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, models.NewResponse(
		http.StatusOK,
		"Profile updated successfully",
		nil,
	))
} 