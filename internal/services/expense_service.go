package services

import (
	"context"
	"time"
	"wanderwallet/internal/models"
	"wanderwallet/internal/repository"
)

type ExpenseService struct {
	repo repository.ExpenseRepositoryInterface
}

func NewExpenseService(repo repository.ExpenseRepositoryInterface) *ExpenseService {
	return &ExpenseService{
		repo: repo,
	}
}

func (s *ExpenseService) CreateExpense(ctx context.Context, expense *models.Expense) error {
	return s.repo.CreateExpense(ctx, expense)
}

func (s *ExpenseService) GetExpenseByID(ctx context.Context, expenseID uint) (*models.Expense, error) {
	return s.repo.GetExpenseByID(ctx, expenseID)
}

func (s *ExpenseService) GetExpensesByUserID(ctx context.Context, id uint) ([]models.Expense, error) {
	return s.repo.GetExpensesByUserID(ctx, id)
}

func (s *ExpenseService) GetExpensesByUserTimeAndCategory(ctx context.Context, userID uint, fromTime *time.Time, toTime *time.Time, categoryID *uint) ([]models.Expense, error) {
	filter := repository.ExpenseFilter{
		UserID:     userID,
		FromTime:   fromTime,
		ToTime:     toTime,
		CategoryID: categoryID,
	}
	return s.repo.GetExpensesByUserTimeAndCategory(ctx, filter)
}

func (s *ExpenseService) ExistsByCategoryID(ctx context.Context, categoryID uint) (bool, error) {
	return s.repo.ExistsByCategoryID(ctx, categoryID)
}

func (s *ExpenseService) UpdateExpense(ctx context.Context, expense *models.Expense) error {
	return s.repo.UpdateExpense(ctx, expense)
}

func (s *ExpenseService) DeleteExpense(ctx context.Context, id uint) error {
	return s.repo.DeleteExpense(ctx, id)
}
