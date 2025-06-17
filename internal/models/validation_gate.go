package models

import (
	"gorm.io/gorm"
)

type ValidationGate struct {
	gorm.Model
	TerminalID uint   `gorm:"index"` // Foreign key to Terminal
	GateID     string `gorm:"type:varchar(50);uniqueIndex;not null"`
	Direction  string `gorm:"type:varchar(10)"` // e.g., "entry", "exit"
	IsActive   bool   `gorm:"default:true"`
}