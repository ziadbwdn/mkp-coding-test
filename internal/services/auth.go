// internal/services/auth.go
package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	
	"pubtrans-eticketing/internal/models"          // Corrected module path
	"pubtrans-eticketing/internal/repositories" // Corrected module path
	customErrors "pubtrans-eticketing/pkg/errors" // Corrected module path
	"pubtrans-eticketing/internal/utils"           // Corrected module path
)

type AuthService interface {
	Login(username, password string) (*models.LoginResponse, error)
}

type authService struct {
	userRepo  repositories.UserRepository
	jwtSecret string
}

func NewAuthService(userRepo repositories.UserRepository, jwtSecret string) AuthService {
	return &authService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (s *authService) Login(username, password string) (*models.LoginResponse, error) {
	// Get user by username
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		if errors.Is(err, customErrors.ErrUserNotFound) {
			return nil, customErrors.NewAppError(
				customErrors.ErrUnauthorized.Code,
				"Invalid username or password",
				err,
			)
		}
		return nil, err 
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, customErrors.NewAppError(
			customErrors.ErrInvalidPassword.Code,
			"Invalid username or password",
			err,
		)
	}

	// Generate JWT token
	// FIX: Cast user.ID to int
	token, err := utils.GenerateJWT(int(user.ID), user.Username, user.Role, s.jwtSecret) // <--- FIX IS HERE
	if err != nil {
		return nil, customErrors.NewAppError(
			customErrors.ErrInternal.Code,
			"Failed to generate authentication token",
			err,
		)
	}

	return &models.LoginResponse{
		Token: token,
		User:  *user,
	}, nil
}