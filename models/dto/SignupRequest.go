package models

type SignupRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	FullName    string `json:"full_name"`
}
