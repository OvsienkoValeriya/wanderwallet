package repository

import (
	"context"
	"wanderwallet/internal/models"

	"gorm.io/gorm"
)

type TravelRepository struct {
	db *gorm.DB
}

func NewTravelRepository(db *gorm.DB) TravelRepositoryInterface {
	return &TravelRepository{db: db}
}

func (r *TravelRepository) CreateTravel(ctx context.Context, travel *models.Travel) error {
	return r.db.WithContext(ctx).Create(travel).Error
}

func (r *TravelRepository) GetTravelByID(ctx context.Context, travelID uint) (*models.Travel, error) {
	var travel models.Travel
	if err := r.db.WithContext(ctx).Where("id = ?", travelID).First(&travel).Error; err != nil {
		return nil, err
	}
	return &travel, nil
}
