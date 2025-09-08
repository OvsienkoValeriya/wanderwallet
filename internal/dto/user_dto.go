package dto

type UserRequest struct {
	Login    string `json:"login" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=6"`
}

type RegisterResponse struct {
	ID    uint   `json:"id"`
	Login string `json:"login"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
