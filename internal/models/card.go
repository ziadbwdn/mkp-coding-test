package models

import (
	"gorm.io/gorm"
	"github.com/shopspring/decimal"
)

type Card struct {
	gorm.Model
	CardNumber   string          `gorm:"type:varchar(50);uniqueIndex;not null"`
	UserID       uint            `gorm:"index"` // Foreign key to User
	Balance      decimal.Decimal `gorm:"type:decimal(10,2);default:0.00"`
	IsActive     bool            `gorm:"default:true"`
	LastUsedAt   *gorm.DeletedAt `gorm:"index"` // Using *gorm.DeletedAt for nullable timestamp or soft delete
	
	Transactions []Transaction `gorm:"foreignKey:CardID"`
}