package services

import (
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

func (s *CategoryService) GetCategoryByID(id uint) (*models.Category, error) {
	return s.repo.GetCategoryByID(id)
}

func (s *CategoryService) GetCategoryByName(name string) (*models.Category, error) {
	return s.repo.GetCategoryByName(name)
}

func (s *CategoryService) GetAllCategories(userID uint) ([]models.Category, error) {
	return s.repo.GetAllCategories(userID)
}

func (s *CategoryService) CreateCategory(category *models.Category) error {
	return s.repo.CreateCategory(category)
}

func (s *CategoryService) DeleteCategory(categoryID uint) error {
	hasExpenses, err := s.expenseRepo.ExistsByCategoryID(categoryID)
	if err != nil {
		return err
	}
	if hasExpenses {
		return ErrCategoryHasLinkedExpenses
	}

	return s.repo.DeleteCategory(categoryID)
}
