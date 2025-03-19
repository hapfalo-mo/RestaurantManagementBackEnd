package handlers

import (
	"RestuarantBackend/interfaces"
	dto "RestuarantBackend/models/dto"
	"net/http"

	"strconv"

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

// Get Paging Booking List Of Detail User
func (b *BookingController) PagingBookingList(c *gin.Context) {
	var request *dto.PagingRequest
	var id int

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	idStr := c.Param("id")
	id, err = strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if b.service == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Internal Error": "Service is not working.."})
	}
	result, err := b.service.PagingBookingList(request, id)
	if err != nil {
		c.JSON(http.StatusFound, gin.H{"error": "Can not show the list"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Data": result})
}

// Get Booking List with out Detail User
func (b *BookingController) PagingAllBookingList(c *gin.Context) {
	var request *dto.PagingRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}
	if b.service == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
	}
	result, err := b.service.PagingAllBookingList(request)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": err})
	}
	c.JSON(http.StatusOK, gin.H{"Data": result})
}
