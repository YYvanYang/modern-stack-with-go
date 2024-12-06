package services

import (
	"gorm.io/gorm"
	"modern-stack/internal/errors"
	"modern-stack/internal/models"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) GetProfile(userID uint) (*models.User, error) {
	var user models.User
	if err := s.db.First(&user, userID).Error; err != nil {
		return nil, errors.ErrUserNotFound
	}
	return &user, nil
}

func (s *UserService) UpdateProfile(userID uint, name string) error {
	result := s.db.Model(&models.User{}).Where("id = ?", userID).Update("name", name)
	if result.Error != nil {
		return errors.ErrServerError
	}
	if result.RowsAffected == 0 {
		return errors.ErrUserNotFound
	}
	return nil
} 