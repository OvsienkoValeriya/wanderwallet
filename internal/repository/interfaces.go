package repository

import (
	"time"
	"wanderwallet/internal/models"
)

type UserRepositoryInterface interface {
	CreateUser(user *models.User) error
	GetByLogin(login string) (*models.User, error)
	GetByID(id uint) (*models.User, error)
	IsLoginExists(login string) (bool, error)
}

type TravelRepositoryInterface interface {
	CreateTravel(travel *models.Travel) error
	GetTravelByID(travelID uint) (*models.Travel, error)
}

type ExpenseRepositoryInterface interface {
	CreateExpense(expense *models.Expense) error
	GetExpensesByUserID(id uint) ([]models.Expense, error)
	GetExpensesByUserTimeAndCategory(userID uint, fromTime *time.Time, toTime *time.Time, categoryID *uint) ([]models.Expense, error)
	GetExpenseByID(expenseID uint) (*models.Expense, error)
	ExistsByCategoryID(categoryID uint) (bool, error)
	UpdateExpense(expense *models.Expense) error
	DeleteExpense(id uint) error
	SumByCategory(userID uint, travelID uint, from, to *time.Time) (map[string]float64, error)
	SumByDay(userID uint, travelID uint, from, to *time.Time) (map[string]float64, error)
	TotalSum(userID uint, travelID uint, from, to *time.Time) (float64, error)
}

type CategoryRepositoryInterface interface {
	GetAllCategories(userID uint) ([]models.Category, error)
	GetCategoryByID(id uint) (*models.Category, error)
	GetCategoryByName(name string) (*models.Category, error)
	CreateCategory(category *models.Category) error
	DeleteCategory(categoryID uint) error
}
