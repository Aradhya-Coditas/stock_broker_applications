package models

type User struct {
	UserName    string `json:"username"`
	PhoneNumber string `json:"phonenumber" validate:"required,min=10,max=15"`
	PanCard     string `json:"pancard" validate:"required,PanCard"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=8,max=32"`
}
