package repositories

import (
	"gorm.io/gorm"
	"github.com/lib/pq" 

	"pubtrans-eticketing/internal/models"
	customErrors "pubtrans-eticketing/pkg/errors" // Alias for clarity
)

type TerminalRepository interface {
	Create(terminal *models.Terminal) error
	GetByCode(code string) (*models.Terminal, error)
	GetAll() ([]models.Terminal, error)
}

type terminalRepository struct {
	db *gorm.DB
}

func NewTerminalRepository(db *gorm.DB) TerminalRepository {
	return &terminalRepository{db: db}
}

func (r *terminalRepository) Create(terminal *models.Terminal) error {
	result := r.db.Create(terminal)

	if result.Error != nil {
		// Check for PostgreSQL unique constraint violation (error code 23505) for 'code'
		if pqErr, ok := result.Error.(*pq.Error); ok && pqErr.Code.Name() == "unique_violation" {
			if pqErr.Constraint == "terminals_code_key" { // Or check against the actual constraint name
				// Use predefined custom error for terminal exists
				return customErrors.ErrTerminalExists
			}
		}
		// Wrap other database errors
		return customErrors.NewAppError(
			customErrors.ErrDatabase.Code,
			"Failed to create terminal in database",
			result.Error,
		)
	}

	return nil
}

func (r *terminalRepository) GetByCode(code string) (*models.Terminal, error) {
	var terminal models.Terminal
	result := r.db.Where("code = ?", code).First(&terminal)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Use predefined custom error for not found
			return nil, customErrors.ErrTerminalNotFound
		}
		// Wrap other database errors
		return nil, customErrors.NewAppError(
			customErrors.ErrDatabase.Code,
			"Failed to retrieve terminal from database",
			result.Error,
		)
	}

	return &terminal, nil
}

func (r *terminalRepository) GetAll() ([]models.Terminal, error) {
	var terminals []models.Terminal
	result := r.db.Where("is_active = ?", true).Order("name").Find(&terminals)

	if result.Error != nil {
		// No need to check for ErrRecordNotFound here with Find
		return nil, customErrors.NewAppError(
			customErrors.ErrDatabase.Code,
			"Failed to retrieve terminals from database",
			result.Error,
		)
	}

	return terminals, nil
}