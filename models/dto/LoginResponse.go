package models

type LoginResponse struct {
	Id          int    `json:"id"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	FullName    string `json:"full_name"`
	Role        int    `json:"role"`
	Point       int    `json:"point"`
}
