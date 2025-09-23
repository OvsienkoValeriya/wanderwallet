package services

import (
	"context"
	"time"
	"wanderwallet/internal/models"
	"wanderwallet/internal/repository"
)

type TravelService struct {
	repo repository.TravelRepositoryInterface
}

func NewTravelService(travelRepo repository.TravelRepositoryInterface) *TravelService {
	return &TravelService{
		repo: travelRepo,
	}
}

func (s *TravelService) CreateTravel(ctx context.Context, userID uint, title string, start, end time.Time) (*models.Travel, error) {
	travel := &models.Travel{UserID: userID, Title: title, StartDate: start, EndDate: end}
	return travel, s.repo.CreateTravel(ctx, travel)
}

func (s *TravelService) GetTravelByID(ctx context.Context, travelID uint) (*models.Travel, error) {
	return s.repo.GetTravelByID(ctx, travelID)
}
