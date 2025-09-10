package dto

type AnalyticsResponse struct {
	Total      float64            `json:"total"`
	ByCategory map[string]float64 `json:"by_category"`
	ByDay      map[string]float64 `json:"by_day"`
}
