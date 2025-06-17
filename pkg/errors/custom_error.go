// pkg/errors/custom_error.go
package errors

import (
	"fmt"
)

// AppError represents a custom error with a specific code and message.
type AppError struct {
	Code    string
	Message string
	Err     error // The underlying error, if any
}

// Error implements the error interface for AppError.
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// Unwrap returns the underlying error. This makes AppError compatible with errors.Is and errors.As.
func (e *AppError) Unwrap() error {
	return e.Err
}

// NewAppError creates a new AppError instance.
func NewAppError(code, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// Predefined error types
var (
	// User related errors
	ErrUserNotFound    = NewAppError("USER_NOT_FOUND", "User not found", nil)
	ErrInvalidPassword = NewAppError("INVALID_PASSWORD", "Invalid password", nil)

	// Terminal related errors
	ErrTerminalExists    = NewAppError("TERMINAL_EXISTS", "Terminal with this code already exists", nil)
	ErrTerminalNotFound  = NewAppError("TERMINAL_NOT_FOUND", "Terminal not found", nil) // NEW: Added Terminal Not Found

	// Authentication/Authorization errors
	ErrInvalidToken   = NewAppError("INVALID_TOKEN", "Invalid or expired token", nil)
	ErrUnauthorized   = NewAppError("UNAUTHORIZED", "Unauthorized access", nil)
	ErrPermissionDenied = NewAppError("PERMISSION_DENIED", "You do not have permission to perform this action", nil) // Common additional auth error

	// Validation errors
	ErrValidationFailed = NewAppError("VALIDATION_FAILED", "Input validation failed", nil)

	// Database/Internal errors
	ErrDatabase = NewAppError("DATABASE_ERROR", "A database error occurred", nil) // NEW: General database error
	ErrInternal = NewAppError("INTERNAL_SERVER_ERROR", "An unexpected internal error occurred", nil) // General catch-all
)