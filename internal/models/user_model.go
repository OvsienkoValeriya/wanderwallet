package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Login    string `gorm:"unique"`
	Password string

	Travels    []Travel   `gorm:"foreignKey:UserID"`
	Categories []Category `gorm:"foreignKey:UserID"`
	Expenses   []Expense  `gorm:"foreignKey:UserID"`
}
