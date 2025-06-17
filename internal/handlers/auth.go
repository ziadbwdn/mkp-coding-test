package handlers

import (
	"net/http"
	"errors"

	"github.com/gin-gonic/gin"

	"pubtrans-eticketing/internal/models"
	"pubtrans-eticketing/internal/services"
	customErrors "pubtrans-eticketing/pkg/errors"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse(
			"Invalid request format",
			customErrors.NewAppError(customErrors.ErrValidationFailed.Code, "Request payload invalid", err),
		))
		return
	}

	loginResponse, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		var appErr *customErrors.AppError
		statusCode := http.StatusInternalServerError 

		if errors.As(err, &appErr) {
			// Map specific AppError codes to HTTP status codes
			switch appErr.Code {
			case customErrors.ErrUserNotFound.Code, customErrors.ErrInvalidPassword.Code:
				statusCode = http.StatusUnauthorized // For login, these generally map to 401
			case customErrors.ErrDatabase.Code:
				statusCode = http.StatusInternalServerError
			case customErrors.ErrInternal.Code:
				statusCode = http.StatusInternalServerError
			default:
				statusCode = http.StatusInternalServerError // Fallback
			}
			c.JSON(statusCode, models.ErrorResponse(appErr.Message, appErr))
		} else {
			// Fallback for non-AppError errors
			c.JSON(statusCode, models.ErrorResponse("Authentication failed", err))
		}
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse(
		"Login successful",
		loginResponse,
	))
}