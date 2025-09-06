package services

import (
	"time"
	"wanderwallet/internal/dto"
	"wanderwallet/internal/models"
)

type UserServiceInterface interface {
	Register(login, password string) (*dto.RegisterResponse, error)
	Login(login, password string) (*dto.LoginResponse, error)
	GetUserByID(id uint) (*models.User, error)
}

type TravelServiceInterface interface {
	CreateTravel(userID uint, title string, start, end time.Time) (*models.Travel, error)
	GetTravelByID(travelID uint) (*models.Travel, error)
}

type ExpenseServiceInterface interface {
	CreateExpense(expense *models.Expense) (*models.Expense, error)
	GetExpensesByUserID(id uint) ([]models.Expense, error)
	UpdateExpense(expense *models.Expense) error
	DeleteExpense(id uint) error
}

type CategoryServiceInterface interface {
	GetCategoryByID(id uint) (*models.Category, error)
	GetCategoryByName(name string) (*models.Category, error)
	CreateCategory(category *models.Category) error
	DeleteCategory(categoryID uint) error
}
