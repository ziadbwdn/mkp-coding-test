// internal/models/terminal.go
package models

import (
	"gorm.io/gorm"
)

type Terminal struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null"`
	Code     string `gorm:"type:varchar(10);uniqueIndex;not null"` // Unique code for the terminal
	Location string `gorm:"type:varchar(255);not null"`
	IsActive bool   `gorm:"default:true"`
}

// CreateTerminalRequest represents the request body for creating a terminal.
type CreateTerminalRequest struct {
	Name     string `json:"name" binding:"required"`
	Code     string `json:"code" binding:"required"`
	Location string `json:"location" binding:"required"`
}