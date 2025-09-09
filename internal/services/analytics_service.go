package services

import (
	"time"
	"wanderwallet/internal/dto"
	"wanderwallet/internal/repository"
)

type AnalyticsService struct {
	repo repository.ExpenseRepositoryInterface
}

func NewAnalyticsService(repo repository.ExpenseRepositoryInterface) *AnalyticsService {
	return &AnalyticsService{
		repo: repo,
	}
}

func (s *AnalyticsService) Aggregate(userID uint, travelID uint, from, to time.Time) (*dto.AnalyticsResponse, error) {
	var fromPtr, toPtr *time.Time
	if !from.IsZero() {
		fromPtr = &from
	}
	if !to.IsZero() {
		toPtr = &to
	}

	total, err := s.repo.TotalSum(userID, travelID, fromPtr, toPtr)
	if err != nil {
		return nil, err
	}

	byCat, err := s.repo.SumByCategory(userID, travelID, fromPtr, toPtr)
	if err != nil {
		return nil, err
	}

	byDay, err := s.repo.SumByDay(userID, travelID, fromPtr, toPtr)
	if err != nil {
		return nil, err
	}

	return &dto.AnalyticsResponse{
		Total:      total,
		ByCategory: byCat,
		ByDay:      byDay,
	}, nil
}
