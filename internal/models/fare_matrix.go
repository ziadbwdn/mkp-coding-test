package models

import (
	"gorm.io/gorm"
	"github.com/shopspring/decimal" // Assuming this is used for Amount
)

type FareMatrix struct {
	gorm.Model
	RouteFrom string          `gorm:"type:varchar(100);not null;uniqueIndex:idx_fare_route"`
	RouteTo   string          `gorm:"type:varchar(100);not null;uniqueIndex:idx_fare_route"`
	Amount    decimal.Decimal `gorm:"type:decimal(10,2);not null"`
	IsActive  bool            `gorm:"default:true"`
}