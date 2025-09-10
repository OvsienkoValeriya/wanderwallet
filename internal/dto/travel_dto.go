package dto

type CreateTravelRequest struct {
	Title     string `json:"title"`
	StartDate string `json:"start_date"` // формат YYYY-MM-DD
	EndDate   string `json:"end_date"`
}
