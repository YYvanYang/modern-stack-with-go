package middleware

import (
	"github.com/gin-gonic/gin"
	"modern-stack/internal/config"
	"time"
	"log"
)

type Middleware struct {
	config *config.Config
}

func NewMiddleware(config *config.Config) *Middleware {
	return &Middleware{config: config}
}

func (m *Middleware) Setup(r *gin.Engine) {
	// Recovery middleware
	r.Use(gin.Recovery())

	// CORS middleware
	r.Use(m.CORS())

	// Logger middleware
	r.Use(m.Logger())
}

func (m *Middleware) CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (m *Middleware) Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		latency := time.Since(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

		if raw != "" {
			path = path + "?" + raw
		}

		log.Printf("| %3d | %13v | %15s | %s | %s | %s |",
			statusCode,
			latency,
			clientIP,
			method,
			path,
			errorMessage,
		)
	}
} 