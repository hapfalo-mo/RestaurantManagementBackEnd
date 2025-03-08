package interfaces

import dto "RestuarantBackend/models/dto"

type BookingInterface interface {
	BookingTable(bookingRequest *dto.BookingRequest) (string, error)
}
