package repositories

import (
	"gorm.io/gorm"
	"github.com/lib/pq" // Import for PostgreSQL-specific error checking

	"pubtrans-eticketing/internal/models"
	customErrors "pubtrans-eticketing/pkg/errors" // Alias for clarity
)

type UserRepository interface {
	GetByUsername(username string) (*models.User, error)
	Create(user *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	result := r.db.Where("username = ? AND is_active = ?", username, true).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, customErrors.ErrUserNotFound
		}
		// Wrap other database errors with custom AppError
		return nil, customErrors.NewAppError(
			customErrors.ErrDatabase.Code,
			"Failed to retrieve user from database",
			result.Error,
		)
	}

	return &user, nil
}

func (r *userRepository) Create(user *models.User) error {
	result := r.db.Create(user)

	if result.Error != nil {
		// Check for PostgreSQL unique constraint violation (error code 23505)
		if pqErr, ok := result.Error.(*pq.Error); ok && pqErr.Code.Name() == "unique_violation" {
			// Determine if it's a username or email conflict
			if pqErr.Constraint == "users_username_key" { // Or check against the actual constraint name
				return customErrors.NewAppError(
					customErrors.ErrValidationFailed.Code, // Or a more specific "USER_USERNAME_EXISTS"
					"Username already exists",
					result.Error,
				)
			}
			if pqErr.Constraint == "users_email_key" { // Or check against the actual constraint name
				return customErrors.NewAppError(
					customErrors.ErrValidationFailed.Code, // Or a more specific "USER_EMAIL_EXISTS"
					"Email already exists",
					result.Error,
				)
			}
		}
		// Wrap other database errors
		return customErrors.NewAppError(
			customErrors.ErrDatabase.Code,
			"Failed to create user in database",
			result.Error,
		)
	}

	return nil
}