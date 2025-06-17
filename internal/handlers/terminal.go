package handlers

import (
	"net/http"
	"errors" 

	"github.com/gin-gonic/gin"

	"pubtrans-eticketing/internal/models"
	"pubtrans-eticketing/internal/services"
	customErrors "pubtrans-eticketing/pkg/errors" 
)

type TerminalHandler struct {
	terminalService services.TerminalService
}

func NewTerminalHandler(terminalService services.TerminalService) *TerminalHandler {
	return &TerminalHandler{
		terminalService: terminalService,
	}
}

func (h *TerminalHandler) CreateTerminal(c *gin.Context) {
	var req models.CreateTerminalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse(
			"Invalid request format",
			customErrors.NewAppError(customErrors.ErrValidationFailed.Code, "Request payload invalid", err),
		))
		return
	}

	terminal, err := h.terminalService.CreateTerminal(&req)
	if err != nil {
		var appErr *customErrors.AppError
		statusCode := http.StatusInternalServerError // Default to internal server error

		if errors.As(err, &appErr) {
			switch appErr.Code {
			case customErrors.ErrTerminalExists.Code:
				statusCode = http.StatusConflict // 409 Conflict 
			case customErrors.ErrValidationFailed.Code:
				statusCode = http.StatusBadRequest // 400 Bad Request 
			case customErrors.ErrDatabase.Code:
				statusCode = http.StatusInternalServerError
			case customErrors.ErrInternal.Code:
				statusCode = http.StatusInternalServerError
			default:
				statusCode = http.StatusInternalServerError // Fallback
			}
			c.JSON(statusCode, models.ErrorResponse(appErr.Message, appErr))
		} else {
			c.JSON(statusCode, models.ErrorResponse("Failed to create terminal", err))
		}
		return
	}

	c.JSON(http.StatusCreated, models.SuccessResponse(
		"Terminal created successfully",
		terminal,
	))
}