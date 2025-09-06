package services

import (
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

func (s *TravelService) CreateTravel(userID uint, title string, start, end time.Time) (*models.Travel, error) {
	travel := &models.Travel{UserID: userID, Title: title, StartDate: start, EndDate: end}
	return travel, s.repo.CreateTravel(travel)
}

func (s *TravelService) GetTravelByID(travelID uint) (*models.Travel, error) {
	return s.repo.GetTravelByID(travelID)
}
