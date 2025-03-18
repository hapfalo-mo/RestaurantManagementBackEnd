package models

import "database/sql"

type User struct {
	Id          int            `json:"id"`
	PhoneNumber string         `json:"phone_number"`
	Password    string         `json:"password"`
	Email       string         `json:"email"`
	FullName    string         `json:"full_name"`
	Role        int            `json:"role"`
	Point       int            `json:"point"`
	CreatedAt   string         `json:"created_at"`
	UpdatedAt   string         `json:"updated_at"`
	DeletedAt   sql.NullString `json:"deleted_at"`
}
