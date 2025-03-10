package service

import (
	"RestuarantBackend/db"
	"RestuarantBackend/interfaces"
	dto "RestuarantBackend/models/dto"
	"errors"
	"time"
)

var _ interfaces.BookingInterface = &BookingService{}

type BookingService struct {
}

func (b BookingService) BookingTable(request *dto.BookingRequest) (message string, err error) {
	// Format Date
	// Check if the date is in the correct format
	t, err := time.Parse("02-01-2006 15:04:05", request.BookingDate)
	if err != nil {
		message = "Date is not in the correct format"
		err = errors.New("Date is not in the correct format")
		return message, err
	}
	_, err = db.DB.Exec(
		"INSERT INTO booking (user_id, customer_name, customer_phone, guest_count, time, note) VALUES (?, ?, ?, ?, ?, ?)",
		request.UserId, request.CustomerName, request.CustomerPhone, request.GuestCount, t.Format("2006-01-02 15:04:05"), request.Description,
	)
	if err != nil {
		message = "Failed to book table"
		err = errors.New("Failed to book table")
		return message, err
	}
	message = "Booking Success. Please check your email or sms for confirmation"
	return message, nil
}

// func (b BookingService) PagingBookingList(request *dto.PagingRequest, userid int) (result []dto.BookingResponse, err error) {
// 	rows, err := db.DB.Query(
// 		`SELECT b.id,b.user_id, b.guest_count, b.time, b.created_at, b.status, b.note, b.customer_name, b.customer_phone, u.full_name, u.phone_number
// 			FROM booking b
// 			JOIN user u
// 			ON b.user_id = u.id
// 			WHERE b.user_id = ?
// 			ORDER BY b.created_at DESC
// 			LIMIT ?
// 			OFFSET ?  `,
// 		userid, request.PageSize, (request.Page-1)*request.PageSize,
// 	)

// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	for rows.Next() {

// 	}
