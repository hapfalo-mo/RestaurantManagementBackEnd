package interfaces

import dto "RestuarantBackend/models/dto"

type BookingInterface interface {
	BookingTable(bookingRequest *dto.BookingRequest) (string, error)
	PagingBookingList(request *dto.PagingRequest, userid int) ([]dto.BookingResponse, error)
	PagingAllBookingList(request *dto.PagingRequest) ([]dto.BookingResponse, error)
}
