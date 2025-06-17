package services

import (
	"errors"
	"strings"

	"pubtrans-eticketing/internal/models"
	"pubtrans-eticketing/internal/repositories"
	customErrors "pubtrans-eticketing/pkg/errors" 
)

// TerminalService defines the interface for terminal-related business logic.
type TerminalService interface {
	CreateTerminal(req *models.CreateTerminalRequest) (*models.Terminal, error)
}

type terminalService struct {
	terminalRepo repositories.TerminalRepository
}

// NewTerminalService creates a new instance of TerminalService.
func NewTerminalService(terminalRepo repositories.TerminalRepository) TerminalService {
	return &terminalService{
		terminalRepo: terminalRepo,
	}
}

// CreateTerminal handles the business logic for creating a new terminal.
func (s *terminalService) CreateTerminal(req *models.CreateTerminalRequest) (*models.Terminal, error) {

	_, err := s.terminalRepo.GetByCode(req.Code) 
	if err == nil { 
		return nil, customErrors.ErrTerminalExists
	}
	
	// If an error occurred, check if it's "not found" (expected) or another DB error
	var appErr *customErrors.AppError
	if errors.As(err, &appErr) && appErr.Code != customErrors.ErrTerminalNotFound.Code {
		return nil, err
	}

	// Create terminal model
	terminal := &models.Terminal{
		Name:     strings.TrimSpace(req.Name),
		Code:     strings.ToUpper(strings.TrimSpace(req.Code)),
		Location: strings.TrimSpace(req.Location),
		IsActive: true,
	}

	if err := s.terminalRepo.Create(terminal); err != nil {
		return nil, err
	}

	return terminal, nil
}