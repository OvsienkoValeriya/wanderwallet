package models

import (
	"time"

	"gorm.io/gorm"
)

type Travel struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	Title     string
	StartDate time.Time
	EndDate   time.Time

	User     User      `gorm:"foreignKey:UserID"`
	Expenses []Expense `gorm:"foreignKey:TravelID"`
}
