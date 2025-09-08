package models

import (
	"time"

	"gorm.io/gorm"
)

type Expense struct {
	gorm.Model
	ID          uint    `gorm:"primaryKey"`
	UserID      uint    `gorm:"not null;index"`
	TravelID    uint    `gorm:"not null;index"`
	CategoryID  uint    `gorm:"index"`
	Amount      float64 `gorm:"not null"`
	Description string
	CreatedAt   time.Time

	User     User     `gorm:"foreignKey:UserID"`
	Travel   Travel   `gorm:"foreignKey:TravelID"`
	Category Category `gorm:"foreignKey:CategoryID"`
}
