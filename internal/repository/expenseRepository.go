package repository

import (
	"time"
	"wanderwallet/internal/models"

	"gorm.io/gorm"
)

type ExpenseRepository struct {
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) ExpenseRepositoryInterface {
	return &ExpenseRepository{db: db}
}

func (r *ExpenseRepository) CreateExpense(expense *models.Expense) error {
	return r.db.Create(expense).Error
}

func (r *ExpenseRepository) GetExpenseByID(expenseID uint) (*models.Expense, error) {
	var expense models.Expense
	err := r.db.Where("id = ?", expenseID).First(&expense).Error
	return &expense, err
}

func (r *ExpenseRepository) GetExpensesByUserID(userID uint) ([]models.Expense, error) {
	var expenses []models.Expense
	err := r.db.Where("user_id = ?", userID).Find(&expenses).Error
	return expenses, err

}

func (r *ExpenseRepository) GetExpensesByUserTimeAndCategory(userID uint, fromTime *time.Time, toTime *time.Time, categoryID *uint) ([]models.Expense, error) {
	var expenses []models.Expense
	query := r.db.Model(&models.Expense{}).Where("user_id = ?", userID)

	if categoryID != nil {
		query = query.Where("category_id = ?", *categoryID)
	}

	if fromTime != nil {
		query = query.Where("created_at >= ?", *fromTime)
	}

	if toTime != nil {
		query = query.Where("created_at <= ?", *toTime)
	}

	err := query.Find(&expenses).Error
	return expenses, err
}

func (r *ExpenseRepository) UpdateExpense(expense *models.Expense) error {
	return r.db.Save(expense).Error

}

func (r *ExpenseRepository) DeleteExpense(id uint) error {
	return r.db.Delete(&models.Expense{}, id).Error
}

func (r *ExpenseRepository) ExistsByCategoryID(categoryID uint) (bool, error) {
	var count int64
	if err := r.db.Model(&models.Expense{}).
		Where("category_id = ?", categoryID).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
