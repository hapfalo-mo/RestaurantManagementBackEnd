package models

type BookingRequest struct {
	UserId        int    `json:"UserId"`
	CustomerName  string `json:"CustomerName"`
	CustomerPhone string `json:"CustomerPhone"`
	GuestCount    int    `json:"GuestCount"`
	BookingDate   string `json:"BookingDate"`
	Description   string `json:"Description"`
}
