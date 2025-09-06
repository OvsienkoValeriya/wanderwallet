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

func TestTravelService_CreateTravel_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTravelRepositoryInterface(ctrl)
	service := NewTravelService(mockRepo)

	userID := uint(1)
	title := "Trip to Paris"
	startDate := time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 3, 25, 0, 0, 0, 0, time.UTC)

	mockRepo.EXPECT().CreateTravel(gomock.Any()).Return(nil).Do(func(travel *models.Travel) {
		travel.ID = 1 // Имитируем автоинкремент
	})

	result, err := service.CreateTravel(userID, title, startDate, endDate)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, userID, result.UserID)
	assert.Equal(t, title, result.Title)
	assert.Equal(t, startDate, result.StartDate)
	assert.Equal(t, endDate, result.EndDate)
}

func TestTravelService_CreateTravel_RepositoryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTravelRepositoryInterface(ctrl)
	service := NewTravelService(mockRepo)

	expectedError := errors.New("database error")
	mockRepo.EXPECT().CreateTravel(gomock.Any()).Return(expectedError)

	result, err := service.CreateTravel(1, "Test", time.Now(), time.Now().Add(24*time.Hour))

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.NotNil(t, result)
}
