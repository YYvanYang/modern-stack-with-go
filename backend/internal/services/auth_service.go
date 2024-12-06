package services

import (
	"context"
	"log"
	"time"
	"strings"

	"gorm.io/gorm"
	customerrors "modern-stack/internal/errors"
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
		return nil, customerrors.ErrInvalidCredentials
	}

	if !user.CheckPassword(password) {
		return nil, customerrors.ErrInvalidCredentials
	}

	return &user, nil
}

// 添加密码验证函数
func validatePassword(password string) error {
	if len(password) < 6 {
		return customerrors.ErrInvalidPassword
	}
	if !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return customerrors.ErrInvalidPassword
	}
	return nil
}

func (s *AuthService) Register(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 在事务开始前验证密码
	if err := validatePassword(user.Password); err != nil {
		return err
	}

	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var existingUser models.User
		if err := tx.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
			return customerrors.ErrEmailExists
		} else if err != gorm.ErrRecordNotFound {
			log.Printf("Error checking existing user: %v", err)
			return customerrors.ErrServerError
		}

		log.Printf("Attempting to create user: %+v", user)
		
		if err := tx.Create(user).Error; err != nil {
			log.Printf("Error creating user: %v", err)
			return customerrors.ErrServerError
		}

		log.Printf("User created successfully: %d", user.ID)
		return nil
	})
} 