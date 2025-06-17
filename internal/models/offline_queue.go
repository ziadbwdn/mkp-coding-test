package models

import (
	"gorm.io/gorm"
	"github.com/shopspring/decimal" // Keep if used for Amount
)

// OfflineQueue stores transactions that happened offline, waiting for synchronization.
type OfflineQueue struct {
	gorm.Model
	TransactionID string          `gorm:"type:varchar(100);uniqueIndex;not null"` // Unique ID for the offline transaction
	CardNumber    string          `gorm:"type:varchar(50);not null"`
	Amount        decimal.Decimal `gorm:"type:decimal(10,2);not null"` // Monetary amount
	TerminalCode  string          `gorm:"type:varchar(10);not null"`
	Timestamp     int64           `gorm:"not null"` // Unix timestamp of the transaction
	Synced        bool            `gorm:"default:false"`
}