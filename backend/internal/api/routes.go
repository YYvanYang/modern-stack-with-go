package api

import (
	"github.com/gin-gonic/gin"
	"modern-stack/internal/config"
	"modern-stack/internal/middleware"
	"modern-stack/internal/services"
)

type Router struct {
	config      *config.Config
	authHandler *AuthHandler
	userHandler *UserHandler
}

func NewRouter(config *config.Config, authService *services.AuthService, userService *services.UserService) *Router {
	return &Router{
		config:      config,
		authHandler: NewAuthHandler(authService, config),
		userHandler: NewUserHandler(userService),
	}
}

func (r *Router) Setup(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")

	// Public routes
	public := v1.Group("")
	{
		public.GET("/health", HealthCheck)
		public.HEAD("/health", HealthCheck)
		auth := public.Group("/auth")
		{
			auth.POST("/login", r.authHandler.Login)
			auth.POST("/register", r.authHandler.Register)
		}
	}

	// Protected routes
	protected := v1.Group("")
	protected.Use(middleware.AuthMiddleware(r.config.JWT.Secret))
	{
		user := protected.Group("/user")
		{
			user.GET("/profile", r.userHandler.GetProfile)
			user.PUT("/profile", r.userHandler.UpdateProfile)
		}
	}
} 