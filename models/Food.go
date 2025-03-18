package models

type Food struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	FoodImage   string `json:"foodImage"`
	IsAvailable bool   `json:"isAvailable"`
	DeletedAt   string `json:"deletedAt"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}
