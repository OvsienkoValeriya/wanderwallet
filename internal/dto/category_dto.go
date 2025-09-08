package dto

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

type CategoryResponse struct {
	ID      string `json:"id"` // string для id
	Name    string `json:"name"`
	Builtin bool   `json:"builtin"`
}
