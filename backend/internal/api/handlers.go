package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"modern-stack/internal/models"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, models.NewResponse(http.StatusOK, "ok", gin.H{
		"version": "1.0.0",
	}))
} 