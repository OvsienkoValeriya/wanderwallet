package services

import (
	"context"
	"time"
	"wanderwallet/internal/dto"
	"wanderwallet/internal/models"
)

type UserServiceInterface interface {
	Register(ctx context.Context, login, password string) (*dto.RegisterResponse, error)
	Login(ctx context.Context, login, password string) (*dto.LoginResponse, error)
	GetUserByID(ctx context.Context, id uint) (*models.User, error)
}

type TravelServiceInterface interface {
	CreateTravel(ctx context.Context, userID uint, title string, start, end time.Time) (*models.Travel, error)
	GetTravelByID(ctx context.Context, travelID uint) (*models.Travel, error)
}

type ExpenseServiceInterface interface {
	CreateExpense(ctx context.Context, expense *models.Expense) (*models.Expense, error)
	GetExpensesByUserID(ctx context.Context, id uint) ([]models.Expense, error)
	UpdateExpense(ctx context.Context, expense *models.Expense) error
	DeleteExpense(ctx context.Context, id uint) error
}

type CategoryServiceInterface interface {
	GetCategoryByID(ctx context.Context, id uint) (*models.Category, error)
	GetCategoryByName(ctx context.Context, name string) (*models.Category, error)
	CreateCategory(ctx context.Context, category *models.Category) error
	DeleteCategory(ctx context.Context, categoryID uint) error
}

type AnalyticsServiceInterfase interface {
	Aggregate(ctx context.Context, userID uint, travelID uint, from time.Time, to time.Time) (*dto.AnalyticsResponse, error)
}
