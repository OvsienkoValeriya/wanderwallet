package services_test

import (
	"errors"
	"testing"
	"wanderwallet/internal/mocks"
	"wanderwallet/internal/models"
	"wanderwallet/internal/repository"
	"wanderwallet/internal/services"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCategoryService_CreateCategory_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategoryRepositoryInterface(ctrl)
	mockExpenseRepo := mocks.NewMockExpenseRepositoryInterface(ctrl)

	svc := services.NewCategoryService(mockRepo, mockExpenseRepo)

	userID := uint(1)
	category := &models.Category{
		UserID:  &userID,
		Name:    "Кофе и чай",
		Builtin: false,
	}

	mockRepo.EXPECT().
		CreateCategory(category).
		Return(nil)

	err := svc.CreateCategory(category)
	assert.NoError(t, err)
}

func TestCategoryService_CreateCategory_NameTaken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategoryRepositoryInterface(ctrl)
	mockExpenseRepo := mocks.NewMockExpenseRepositoryInterface(ctrl)

	svc := services.NewCategoryService(mockRepo, mockExpenseRepo)

	userID := uint(1)
	category := &models.Category{
		UserID:  &userID,
		Name:    "Кофе и чай",
		Builtin: false,
	}

	mockRepo.EXPECT().
		CreateCategory(category).
		Return(repository.ErrCategoryExists)

	err := svc.CreateCategory(category)
	assert.ErrorIs(t, err, repository.ErrCategoryExists)
}

func TestCategoryService_DeleteCategory_NoExpenses_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategoryRepositoryInterface(ctrl)
	mockExpenseRepo := mocks.NewMockExpenseRepositoryInterface(ctrl)

	svc := services.NewCategoryService(mockRepo, mockExpenseRepo)

	categoryID := uint(5)

	mockExpenseRepo.EXPECT().
		ExistsByCategoryID(categoryID).
		Return(false, nil)

	mockRepo.EXPECT().
		DeleteCategory(categoryID).
		Return(nil)

	err := svc.DeleteCategory(categoryID)
	assert.NoError(t, err)
}

func TestCategoryService_DeleteCategory_HasExpenses_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategoryRepositoryInterface(ctrl)
	mockExpenseRepo := mocks.NewMockExpenseRepositoryInterface(ctrl)

	svc := services.NewCategoryService(mockRepo, mockExpenseRepo)

	categoryID := uint(5)

	mockExpenseRepo.EXPECT().
		ExistsByCategoryID(categoryID).
		Return(true, nil)

	err := svc.DeleteCategory(categoryID)
	assert.ErrorIs(t, err, services.ErrCategoryHasLinkedExpenses)
}

func TestCategoryService_DeleteCategory_RepoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategoryRepositoryInterface(ctrl)
	mockExpenseRepo := mocks.NewMockExpenseRepositoryInterface(ctrl)

	svc := services.NewCategoryService(mockRepo, mockExpenseRepo)

	categoryID := uint(5)
	mockExpenseRepo.EXPECT().
		ExistsByCategoryID(categoryID).
		Return(false, nil)

	mockRepo.EXPECT().
		DeleteCategory(categoryID).
		Return(errors.New("db error"))

	err := svc.DeleteCategory(categoryID)
	assert.Error(t, err)
}

func TestCategoryService_GetCategoryByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategoryRepositoryInterface(ctrl)
	mockExpenseRepo := mocks.NewMockExpenseRepositoryInterface(ctrl)
	svc := services.NewCategoryService(mockRepo, mockExpenseRepo)

	category := &models.Category{ID: 1, Name: "Test"}
	mockRepo.EXPECT().GetCategoryByID(uint(1)).Return(category, nil)

	res, err := svc.GetCategoryByID(1)
	assert.NoError(t, err)
	assert.Equal(t, category, res)
}

func TestCategoryService_GetAllCategories(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategoryRepositoryInterface(ctrl)
	mockExpenseRepo := mocks.NewMockExpenseRepositoryInterface(ctrl)
	svc := services.NewCategoryService(mockRepo, mockExpenseRepo)

	cats := []models.Category{{ID: 1, Name: "A"}, {ID: 2, Name: "B"}}
	mockRepo.EXPECT().GetAllCategories(uint(1)).Return(cats, nil)

	res, err := svc.GetAllCategories(1)
	assert.NoError(t, err)
	assert.Equal(t, cats, res)
}

func TestCategoryService_GetCategoryByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategoryRepositoryInterface(ctrl)
	mockExpenseRepo := mocks.NewMockExpenseRepositoryInterface(ctrl)
	svc := services.NewCategoryService(mockRepo, mockExpenseRepo)

	cat := &models.Category{ID: 1, Name: "Food"}
	mockRepo.EXPECT().GetCategoryByName("Food").Return(cat, nil)

	res, err := svc.GetCategoryByName("Food")
	assert.NoError(t, err)
	assert.Equal(t, cat, res)
}
