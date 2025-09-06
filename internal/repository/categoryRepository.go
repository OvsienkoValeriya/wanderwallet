package repository

import (
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

func (r *CategoryRepository) GetCategoryByID(id uint) (*models.Category, error) {
	var category models.Category
	err := r.db.Where("id = ?", id).First(&category).Error
	return &category, err
}

func (r *CategoryRepository) GetCategoryByName(name string) (*models.Category, error) {
	var category models.Category
	if err := r.db.Where("name = ?", name).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) GetAllCategories(userID uint) ([]models.Category, error) {
	var categories []models.Category
	if err := r.db.Where(" builtin = true OR user_id = ?", userID).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoryRepository) CreateCategory(category *models.Category) error {
	var existing models.Category
	if err := r.db.Where("user_id = ? AND name = ?", category.UserID, category.Name).First(&existing).Error; err == nil {
		return ErrCategoryExists
	}
	return r.db.Create(category).Error
}

func (r *CategoryRepository) DeleteCategory(categoryID uint) error {
	return r.db.Delete(&models.Category{}, categoryID).Error
}
