package services

import (
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

func (s *ExpenseService) CreateExpense(expense *models.Expense) error {
	return s.repo.CreateExpense(expense)
}

func (s *ExpenseService) GetExpenseByID(expenseID uint) (*models.Expense, error) {
	return s.repo.GetExpenseByID(expenseID)
}

func (s *ExpenseService) GetExpensesByUserID(id uint) ([]models.Expense, error) {
	return s.repo.GetExpensesByUserID(id)
}

func (s *ExpenseService) GetExpensesByUserTimeAndCategory(userID uint, fromTime *time.Time, toTime *time.Time, categoryID *uint) ([]models.Expense, error) {
	return s.repo.GetExpensesByUserTimeAndCategory(userID, fromTime, toTime, categoryID)
}

func (s *ExpenseService) ExistsByCategoryID(categoryID uint) (bool, error) {
	return s.repo.ExistsByCategoryID(categoryID)
}

func (s *ExpenseService) UpdateExpense(expense *models.Expense) error {
	return s.repo.UpdateExpense(expense)
}

func (s *ExpenseService) DeleteExpense(id uint) error {
	return s.repo.DeleteExpense(id)
}
