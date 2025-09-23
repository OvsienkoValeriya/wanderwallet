package repository

import (
	"context"
	"errors"
	"wanderwallet/internal/models"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

var ErrCategoryExists = errors.New("category already exists")

func NewCategoryRepository(db *gorm.DB) CategoryRepositoryInterface {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetCategoryByID(ctx context.Context, id uint) (*models.Category, error) {
	var category models.Category
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&category).Error
	return &category, err
}

func (r *CategoryRepository) GetCategoryByName(ctx context.Context, name string) (*models.Category, error) {
	var category models.Category
	if err := r.db.WithContext(ctx).Where("name = ?", name).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) GetAllCategories(ctx context.Context, userID uint) ([]models.Category, error) {
	var categories []models.Category
	if err := r.db.WithContext(ctx).Where(" builtin = true OR user_id = ?", userID).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoryRepository) CreateCategory(ctx context.Context, category *models.Category) error {
	var existing models.Category
	if err := r.db.WithContext(ctx).Where("user_id = ? AND name = ?", category.UserID, category.Name).First(&existing).Error; err == nil {
		return ErrCategoryExists
	}
	return r.db.Create(category).Error
}

func (r *CategoryRepository) DeleteCategory(ctx context.Context, categoryID uint) error {
	return r.db.WithContext(ctx).Delete(&models.Category{}, categoryID).Error
}
