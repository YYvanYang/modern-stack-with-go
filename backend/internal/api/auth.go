package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"modern-stack/internal/errors"
	"modern-stack/internal/models"
	"modern-stack/internal/services"
	"modern-stack/internal/config"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
}

type AuthHandler struct {
	authService *services.AuthService
	config     *config.Config
}

func NewAuthHandler(authService *services.AuthService, config *config.Config) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		config:     config,
	}
}

func (h *AuthHandler) generateToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(h.config.JWT.TokenDuration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(h.config.JWT.Secret))
}

func (h *AuthHandler) handleError(c *gin.Context, err error) {
	switch err {
	case errors.ErrInvalidCredentials:
		c.JSON(http.StatusUnauthorized, models.NewResponse(
			http.StatusUnauthorized,
			err.Error(),
			nil,
		))
	case errors.ErrEmailExists:
		c.JSON(http.StatusConflict, models.NewResponse(
			http.StatusConflict,
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

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.NewResponse(
			http.StatusBadRequest,
			err.Error(),
			nil,
		))
		return
	}

	user, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		h.handleError(c, err)
		return
	}

	token, err := h.generateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewResponse(
			http.StatusInternalServerError,
			"Failed to generate token",
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, models.NewResponse(http.StatusOK, "Login successful", gin.H{
		"token": token,
		"user":  user,
	}))
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.NewResponse(
			http.StatusBadRequest,
			"请检查输入信息格式是否正确",
			nil,
		))
		return
	}

	user := &models.User{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
		Role:     "user",
	}

	if err := h.authService.Register(user); err != nil {
		switch err {
		case errors.ErrEmailExists:
			c.JSON(http.StatusConflict, models.NewResponse(
				http.StatusConflict,
				"该邮箱已被注册",
				nil,
			))
		case errors.ErrInvalidPassword:
			c.JSON(http.StatusBadRequest, models.NewResponse(
				http.StatusBadRequest,
				"密码必须包含至少一个大写字母",
				nil,
			))
		default:
			c.JSON(http.StatusInternalServerError, models.NewResponse(
				http.StatusInternalServerError,
				"服务器错误，请稍后重试",
				nil,
			))
		}
		return
	}

	token, err := h.generateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewResponse(
			http.StatusInternalServerError,
			"Failed to generate token",
			nil,
		))
		return
	}

	c.JSON(http.StatusCreated, models.NewResponse(http.StatusCreated, "Registration successful", gin.H{
		"token": token,
		"user":  user,
	}))
} 