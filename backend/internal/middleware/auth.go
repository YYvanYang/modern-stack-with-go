package middleware

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"modern-stack/internal/models"
)

var (
	ErrMissingHeader = errors.New("Authorization header is required")
	ErrInvalidFormat = errors.New("Invalid authorization header format")
	ErrInvalidToken  = errors.New("Invalid token")
)

func AuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			log.Printf("Missing authorization header")
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.NewResponse(
				http.StatusUnauthorized,
				ErrMissingHeader.Error(),
				nil,
			))
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			log.Printf("Invalid authorization header format: %s", authHeader)
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.NewResponse(
				http.StatusUnauthorized,
				ErrInvalidFormat.Error(),
				nil,
			))
			return
		}

		token := parts[1]
		claims := jwt.MapClaims{}

		parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			// 验证签名算法
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, ErrInvalidToken
			}
			return []byte(secret), nil
		})

		if err != nil || !parsedToken.Valid {
			log.Printf("Token validation failed: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.NewResponse(
				http.StatusUnauthorized,
				ErrInvalidToken.Error(),
				nil,
			))
			return
		}

		// 类型断言确保安全
		userID, ok := claims["user_id"].(float64)
		if !ok {
			log.Printf("Invalid user ID type in token claims")
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.NewResponse(
				http.StatusUnauthorized,
				"Invalid user ID in token",
				nil,
			))
			return
		}

		role, ok := claims["role"].(string)
		if !ok {
			log.Printf("Invalid role type in token claims")
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.NewResponse(
				http.StatusUnauthorized,
				"Invalid role in token",
				nil,
			))
			return
		}

		if claims["exp"] != nil {
			if exp, ok := claims["exp"].(float64); ok {
				if time.Now().Unix() > int64(exp) {
					log.Printf("Token expired")
					c.AbortWithStatusJSON(http.StatusUnauthorized, models.NewResponse(
						http.StatusUnauthorized,
						"Token expired",
						nil,
					))
					return
				}
			}
		}

		c.Set("user_id", uint(userID))
		c.Set("user_role", role)
		c.Next()
	}
} 