package handlers

import (
	"RestuarantBackend/interfaces"
	dto "RestuarantBackend/models/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookingController struct {
	service interfaces.BookingInterface
}

// Constructor for denpendency injection
func NewBookingController(service interfaces.BookingInterface) *BookingController {
	if service == nil {
		panic("NewBookingController service is nil")
	}
	return &BookingController{service: service}
}

// Booking Table
func (b *BookingController) BookingTable(c *gin.Context) {
	var BookingRequest *dto.BookingRequest
	err := c.ShouldBindJSON(&BookingRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if b.service == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Service is not initialized"})
		return
	}
	result, err := b.service.BookingTable(BookingRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": result})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Please check confirmation from email or sms for your booking"})
}
