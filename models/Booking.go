package models

type Booking struct {
	id               int    `json:"id"`
	user_id          int    `json:"UserId"`
	guestCount       int    `json:"GuestCount"`
	bookingDate      string `json:"BookingDate"`
	bookingCreatedAt string `json:"BookingCreatedAt"`
	bookingUpdatedAt string `json:"BookingUpdatedAt"`
	status           int    `json:"Status"`
	description      string `json:"Description"`
}
