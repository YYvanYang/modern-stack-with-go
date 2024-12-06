package services

import (
	"context"
	"time"
	"gorm.io/gorm"
	"modern-stack/internal/errors"
	"modern-stack/internal/models"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) Login(email, password string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	if err := s.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.ErrInvalidCredentials
	}

	if !user.CheckPassword(password) {
		return nil, errors.ErrInvalidCredentials
	}

	return &user, nil
}

func (s *AuthService) Register(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var existingUser models.User
		if err := tx.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
			return errors.ErrEmailExists
		}

		if err := tx.Create(user).Error; err != nil {
			return errors.ErrServerError
		}

		return nil
	})
} 