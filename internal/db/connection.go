// internal/db/connection.go (updated NewConnection signature and usage)
package db

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"

	"pubtrans-eticketing/internal/models"
)

// NewConnection initializes and returns a new GORM database connection.
// It now accepts maxOpenConns and maxIdleConns for connection pool configuration.
func NewConnection(databaseURL string, maxOpenConns, maxIdleConns int) (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent, // Change to logger.Info for full query logs in debug
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database with GORM: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	// Use values from config for connection pool settings
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	sqlDB.SetConnMaxIdleTime(1 * time.Minute)

	if err := sqlDB.Ping(); err != nil {
		// Fix: Removed the extra ", err" at the end of the format string
		return nil, fmt.Errorf("failed to ping database after GORM connection: %w", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Terminal{},
		&models.Card{},
		&models.FareMatrix{},
		&models.Transaction{},
		&models.ValidationGate{},
		&models.OfflineQueue{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to auto migrate database: %w", err)
	}

	return db, nil
}