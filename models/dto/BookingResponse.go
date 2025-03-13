package models

type BookingResponse struct {
	Id            int    `json:"id"`
	UserId        int    `json:"UserId"`
	UserName      string `json:"UserName"`
	UserPhone     string `json:"UserPhone"`
	CustomerName  string `json:"CustomerName"`
	CustomerPhone string `json:"CustomerPhone"`
	GuestCount    int    `json:"GuestCount"`
	BookingDate   string `json:"BookingDate"`
	Description   string `json:"Description"`
	Status        string `json:"Status"`
	CreatedAt     string `json:"CreatedAt"`
}
