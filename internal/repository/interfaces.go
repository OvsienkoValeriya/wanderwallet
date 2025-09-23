package repository

import (
	"context"
	"time"
	"wanderwallet/internal/models"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetByLogin(ctx context.Context, login string) (*models.User, error)
	GetByID(ctx context.Context, id uint) (*models.User, error)
	IsLoginExists(ctx context.Context, login string) (bool, error)
}

type TravelRepositoryInterface interface {
	CreateTravel(ctx context.Context, travel *models.Travel) error
	GetTravelByID(ctx context.Context, travelID uint) (*models.Travel, error)
}

type ExpenseRepositoryInterface interface {
	CreateExpense(ctx context.Context, expense *models.Expense) error
	GetExpensesByUserID(ctx context.Context, id uint) ([]models.Expense, error)
	GetExpensesByUserTimeAndCategory(ctx context.Context, filter ExpenseFilter) ([]models.Expense, error)
	GetExpenseByID(ctx context.Context, expenseID uint) (*models.Expense, error)
	ExistsByCategoryID(ctx context.Context, categoryID uint) (bool, error)
	UpdateExpense(ctx context.Context, expense *models.Expense) error
	DeleteExpense(ctx context.Context, id uint) error
	SumByCategory(ctx context.Context, userID uint, travelID uint, from, to *time.Time) (map[string]float64, error)
	SumByDay(ctx context.Context, userID uint, travelID uint, from, to *time.Time) (map[string]float64, error)
	TotalSum(ctx context.Context, userID uint, travelID uint, from, to *time.Time) (float64, error)
}

type CategoryRepositoryInterface interface {
	GetAllCategories(ctx context.Context, userID uint) ([]models.Category, error)
	GetCategoryByID(ctx context.Context, id uint) (*models.Category, error)
	GetCategoryByName(ctx context.Context, name string) (*models.Category, error)
	CreateCategory(ctx context.Context, category *models.Category) error
	DeleteCategory(ctx context.Context, categoryID uint) error
}
