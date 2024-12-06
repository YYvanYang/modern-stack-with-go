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

	// 验证密码长度
	if u.Password != "" && len(u.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	// 验证名称
	if u.Name == "" {
		return errors.New("name is required")
	}

	// 添加邮箱格式验证
	if !isValidEmail(u.Email) {
		return errors.New("invalid email format")
	}

	// 验证密码
	if u.Password != "" {
		if err := isValidPassword(u.Password); err != nil {
			return err
		}
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

// 添加密码验证函数
func isValidPassword(password string) error {
	if len(password) < 6 {
		return errors.New("password must be at least 6 characters")
	}
	// 检查是否包含数字
	if !strings.ContainsAny(password, "0123456789") {
		return errors.New("password must contain at least one number")
	}
	// 检查是否包含大写字母
	if !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return errors.New("password must contain at least one uppercase letter")
	}
	return nil
} 