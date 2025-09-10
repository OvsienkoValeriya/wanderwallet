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

func (r *ExpenseRepository) SumByCategory(userID uint, travelID uint, from, to *time.Time) (map[string]float64, error) {
	var results []struct {
		Category string
		Amount   float64
	}
	query := r.db.Table("expenses").
		Select("categories.name as category, SUM(expenses.amount) as amount").
		Joins("LEFT JOIN categories ON expenses.category_id = categories.id").
		Where("expenses.user_id = ? AND expenses.travel_id = ? AND expenses.deleted_at IS NULL", userID, travelID).
		Group("categories.name")
	if from != nil {
		query = query.Where("expenses.created_at >= ?", *from)
	}
	if to != nil {
		query = query.Where("expenses.created_at <= ?", *to)
	}
	if err := query.Scan(&results).Error; err != nil {
		return nil, err
	}
	res := make(map[string]float64)
	for _, r := range results {
		res[r.Category] = r.Amount
	}
	return res, nil
}

func (r *ExpenseRepository) SumByDay(userID uint, travelID uint, from, to *time.Time) (map[string]float64, error) {
	var results []struct {
		Day    string
		Amount float64
	}
	query := r.db.Table("expenses").
		Select("DATE(expenses.created_at) as day, SUM(expenses.amount) as amount").
		Where("expenses.user_id = ? AND expenses.travel_id = ? AND expenses.deleted_at IS NULL", userID, travelID).
		Group("day")
	if from != nil {
		query = query.Where("expenses.created_at >= ?", *from)
	}
	if to != nil {
		query = query.Where("expenses.created_at <= ?", *to)
	}
	if err := query.Scan(&results).Error; err != nil {
		return nil, err
	}
	res := make(map[string]float64)
	for _, r := range results {
		res[r.Day] = r.Amount
	}
	return res, nil
}

func (r *ExpenseRepository) TotalSum(userID uint, travelID uint, from, to *time.Time) (float64, error) {
	var sum float64
	query := r.db.Model(&models.Expense{}).
		Select("SUM(amount)").
		Where("user_id = ? AND travel_id = ? AND deleted_at IS NULL", userID, travelID)
	if from != nil {
		query = query.Where("created_at >= ?", *from)
	}
	if to != nil {
		query = query.Where("created_at <= ?", *to)
	}
	err := query.Scan(&sum).Error
	return sum, err
}
