package models

type SignInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type SignInResponse struct {
    Message string `json:"message,omitempty"`
    Error   string `json:"error,omitempty"`
}