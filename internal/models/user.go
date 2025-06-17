package models

import (
	"gorm.io/gorm" // Import GORM
)

type User struct {
	gorm.Model

	Username       string `gorm:"type:varchar(100);uniqueIndex;not null" json:"username"` 
	Email          string `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	PasswordHash   string `gorm:"type:varchar(255);not null" json:"-"`
	Role           string `gorm:"type:varchar(50);default:'user';not null" json:"role"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}