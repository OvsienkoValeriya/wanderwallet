package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"unique;not null"`
	UserID  *uint  // null → встроенная, не null → пользовательская
	Builtin bool   `gorm:"default:false"`

	User     *User     `gorm:"foreignKey:UserID"`
	Expenses []Expense `gorm:"foreignKey:CategoryID"`
}
