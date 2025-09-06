package services

import (
	"errors"
	"testing"
	"time"
	"wanderwallet/internal/mocks"
	"wanderwallet/internal/models"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestExpenseService_CreateExpense(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockExpenseRepositoryInterface(ctrl)
	service := NewExpenseService(mockRepo)

	expense := &models.Expense{ID: 1, Amount: 100}

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().CreateExpense(expense).Return(nil)

		err := service.CreateExpense(expense)
		assert.NoError(t, err)
	})

	t.Run("repo error", func(t *testing.T) {
		mockRepo.EXPECT().CreateExpense(expense).Return(errors.New("db error"))

		err := service.CreateExpense(expense)
		assert.EqualError(t, err, "db error")
	})
}

func TestExpenseService_GetExpensesByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockExpenseRepositoryInterface(ctrl)
	service := NewExpenseService(mockRepo)

	expected := []models.Expense{{ID: 1, Amount: 200}}

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().GetExpensesByUserID(uint(1)).Return(expected, nil)

		expenses, err := service.GetExpensesByUserID(1)
		assert.NoError(t, err)
		assert.Equal(t, expected, expenses)
	})

	t.Run("repo error", func(t *testing.T) {
		mockRepo.EXPECT().GetExpensesByUserID(uint(1)).Return(nil, errors.New("db error"))

		expenses, err := service.GetExpensesByUserID(1)
		assert.Error(t, err)
		assert.Nil(t, expenses)
	})
}

func TestExpenseService_GetExpensesByUserTimeAndCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockExpenseRepositoryInterface(ctrl)
	service := NewExpenseService(mockRepo)

	userID := uint(1)
	from := time.Now().Add(-24 * time.Hour)
	to := time.Now()
	categoryID := uint(2)

	expected := []models.Expense{{ID: 10, Amount: 500}}

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().
			GetExpensesByUserTimeAndCategory(userID, &from, &to, &categoryID).
			Return(expected, nil)

		expenses, err := service.GetExpensesByUserTimeAndCategory(userID, &from, &to, &categoryID)
		assert.NoError(t, err)
		assert.Equal(t, expected, expenses)
	})

	t.Run("repo error", func(t *testing.T) {
		mockRepo.EXPECT().
			GetExpensesByUserTimeAndCategory(userID, &from, &to, &categoryID).
			Return(nil, errors.New("db error"))

		expenses, err := service.GetExpensesByUserTimeAndCategory(userID, &from, &to, &categoryID)
		assert.Error(t, err)
		assert.Nil(t, expenses)
	})
}

func TestExpenseService_UpdateExpense(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockExpenseRepositoryInterface(ctrl)
	service := NewExpenseService(mockRepo)

	expense := &models.Expense{ID: 1, Amount: 300}

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().UpdateExpense(expense).Return(nil)

		err := service.UpdateExpense(expense)
		assert.NoError(t, err)
	})

	t.Run("repo error", func(t *testing.T) {
		mockRepo.EXPECT().UpdateExpense(expense).Return(errors.New("db error"))

		err := service.UpdateExpense(expense)
		assert.EqualError(t, err, "db error")
	})
}

func TestExpenseService_DeleteExpense(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockExpenseRepositoryInterface(ctrl)
	service := NewExpenseService(mockRepo)

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().DeleteExpense(uint(1)).Return(nil)

		err := service.DeleteExpense(1)
		assert.NoError(t, err)
	})

	t.Run("repo error", func(t *testing.T) {
		mockRepo.EXPECT().DeleteExpense(uint(1)).Return(errors.New("db error"))

		err := service.DeleteExpense(1)
		assert.EqualError(t, err, "db error")
	})
}
