package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"modern-stack/internal/errors"
	"modern-stack/internal/models"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, models.NewResponse(http.StatusOK, "ok", gin.H{
		"version": "1.0.0",
	}))
}

func GetProfile(c *gin.Context) {
	// TODO: 实现获取用户资料逻辑
	c.JSON(http.StatusNotImplemented, errors.NewAPIError(http.StatusNotImplemented, "Not implemented"))
}

func UpdateProfile(c *gin.Context) {
	// TODO: 实现更新用户资料逻辑
	c.JSON(http.StatusNotImplemented, errors.NewAPIError(http.StatusNotImplemented, "Not implemented"))
} 