package models

import (
	"time"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"errors"
)

type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Email     string         `gorm:"uniqueIndex;size:255" json:"email"`
	Password  string         `gorm:"size:255" json:"-"`
	Name      string         `gorm:"size:255" json:"name"`
	Role      string         `gorm:"size:50;default:'user'" json:"role"`
}

// BeforeSave 在保存前加密密码
func (u *User) BeforeSave(tx *gorm.DB) error {
	// 验证邮箱格式
	if u.Email == "" {
		return errors.New("email is required")
	}

	// 验证名称
	if u.Name == "" {
		return errors.New("name is required")
	}

	// 添加邮箱格式验证
	if !isValidEmail(u.Email) {
		return errors.New("invalid email format")
	}

	// 加密密码
	if u.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}

	return nil
}

// CheckPassword 验证密码
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}

// BeforeCreate 在创建前设置默认值
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.Role == "" {
		u.Role = "user"
	}
	return u.BeforeSave(tx)
}

// 添加邮箱格式验证
func isValidEmail(email string) bool {
	// 简单的邮箱格式验证
	return strings.Contains(email, "@") && strings.Contains(email, ".")
} 