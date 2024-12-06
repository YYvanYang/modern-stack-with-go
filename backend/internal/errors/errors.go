package errors

import "errors"

// 定义业务错误
var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserNotFound      = errors.New("user not found")
	ErrEmailExists       = errors.New("email already exists")
	ErrInvalidToken      = errors.New("invalid token")
	ErrServerError       = errors.New("internal server error")
)

// APIError 定义API错误响应结构
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewAPIError 创建新的API错误
func NewAPIError(code int, message string) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
	}
} 