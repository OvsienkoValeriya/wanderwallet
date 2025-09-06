package repository

import (
	"wanderwallet/internal/models"

	"gorm.io/gorm"
)

type TravelRepository struct {
	db *gorm.DB
}

func NewTravelRepository(db *gorm.DB) TravelRepositoryInterface {
	return &TravelRepository{db: db}
}

func (r *TravelRepository) CreateTravel(travel *models.Travel) error {
	return r.db.Create(travel).Error
}

func (r *TravelRepository) GetTravelByID(travelID uint) (*models.Travel, error) {
	var travel models.Travel
	if err := r.db.Where("id = ?", travelID).First(&travel).Error; err != nil {
		return nil, err
	}
	return &travel, nil
}
