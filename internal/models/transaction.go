// internal/models/terminal.go
package models

import (
	"gorm.io/gorm"
	"github.com/shopspring/decimal"
)

type Transaction struct {
	gorm.Model
	CardID       uint            `gorm:"index"` // Foreign key to Card
	Amount       decimal.Decimal `gorm:"type:decimal(10,2);not null"`
	TerminalCode string          `gorm:"type:varchar(10);not null"` // Code of the terminal where transaction occurred
	Type         string          `gorm:"type:varchar(20);not null"` // e.g., "debit", "credit"
	Status       string          `gorm:"type:varchar(20);not null"` // e.g., "success", "failed"
	ReferenceID  string          `gorm:"type:varchar(100);uniqueIndex"` // Optional: For linking to external systems
}