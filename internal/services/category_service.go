package services

import (
	"context"
	"errors"
	"wanderwallet/internal/models"
	"wanderwallet/internal/repository"
)

type CategoryService struct {
	repo        repository.CategoryRepositoryInterface
	expenseRepo repository.ExpenseRepositoryInterface
}

var (
	ErrCategoryHasLinkedExpenses = errors.New("category has linked expenses")
)

func NewCategoryService(repo repository.CategoryRepositoryInterface, expenseRepo repository.ExpenseRepositoryInterface) *CategoryService {
	return &CategoryService{
		repo:        repo,
		expenseRepo: expenseRepo,
	}
}

func (s *CategoryService) GetCategoryByID(ctx context.Context, id uint) (*models.Category, error) {
	return s.repo.GetCategoryByID(ctx, id)
}

func (s *CategoryService) GetCategoryByName(ctx context.Context, name string) (*models.Category, error) {
	return s.repo.GetCategoryByName(ctx, name)
}

func (s *CategoryService) GetAllCategories(ctx context.Context, userID uint) ([]models.Category, error) {
	return s.repo.GetAllCategories(ctx, userID)
}

func (s *CategoryService) CreateCategory(ctx context.Context, category *models.Category) error {
	return s.repo.CreateCategory(ctx, category)
}

func (s *CategoryService) DeleteCategory(ctx context.Context, categoryID uint) error {
	hasExpenses, err := s.expenseRepo.ExistsByCategoryID(ctx, categoryID)
	if err != nil {
		return err
	}
	if hasExpenses {
		return ErrCategoryHasLinkedExpenses
	}

	return s.repo.DeleteCategory(ctx, categoryID)
}
