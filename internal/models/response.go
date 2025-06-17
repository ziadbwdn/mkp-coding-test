package models

import (
	"errors" 
	customErrors "pubtrans-eticketing/pkg/errors" 
)

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"` 
}

func SuccessResponse(message string, data interface{}) APIResponse {
	return APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func ErrorResponse(message string, err error) APIResponse {
	resp := APIResponse{
		Success: false,
		Message: message,
	}

	// Try to unwrap and cast to AppError
	var appErr *customErrors.AppError
	if errors.As(err, &appErr) {

		resp.Message = appErr.Message
		resp.Error = appErr.Code

	} else {

		resp.Error = err.Error() // Fallback to standard error string
	}
	return resp
}