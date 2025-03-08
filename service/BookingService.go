package service

import (
	"RestuarantBackend/db"
	"RestuarantBackend/interfaces"
	dto "RestuarantBackend/models/dto"
	"errors"
)

var _ interfaces.BookingInterface = &BookingService{}

type BookingService struct {
}

func (b BookingService) BookingTable(request *dto.BookingRequest) (message string, err error) {
	_, err = db.DB.Exec(
		"INSERT INTO booking (user_id, customer_name, customer_phone, guest_count, time, note) VALUES (?, ?, ?, ?, ?, ?)",
		request.User_id, request.CustomerName, request.CustomerPhone, request.GuestCount, request.BookingDate, request.Description,
	)
	if err != nil {
		message = "Failed to book table"
		err = errors.New("Failed to book table")
		return message, err
	}
	message = "Booking Success. Please check your email or sms for confirmation"
	return message, nil
}
