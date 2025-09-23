package dto

type CreateExpenseRequest struct {
	TravelID uint    `json:"travel_id" binding:"required"`
	Category string  `json:"category" binding:"required"`
	Amount   float64 `json:"amount" binding:"required"`
	Date     string  `json:"date" binding:"required"`
	Comment  string  `json:"comment"`
}

type GetUsersExpenseRequest struct {
	Category string `form:"category"`
	From     string `form:"from"` // формат YYYY-MM-DD
	To       string `form:"to"`
}

type ExpenseResponse struct {
	ID       string  `json:"id"`
	Category string  `json:"category"`
	Amount   float64 `json:"amount"`
	Date     string  `json:"date"`
	Comment  string  `json:"comment"`
}

type UpdateExpenseRequest struct {
	Category string  `json:"category"`
	Date     string  `json:"date"` // формат YYYY-MM-DD
	Amount   float64 `json:"amount"`
	Comment  string  `json:"comment"`
}
