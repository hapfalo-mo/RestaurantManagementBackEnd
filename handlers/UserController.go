package handlers

import (
	"RestuarantBackend/interfaces"
	"net/http"

	dto "RestuarantBackend/models/dto"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service interfaces.UserInterface
}

// Constructor for denpendency injection
func NewUserController(service interfaces.UserInterface) *UserController {
	if service == nil {
		panic("UserController NewUserController service is nil")
	}
	return &UserController{service: service}
}

// User Sign Up
func (u *UserController) Register(c *gin.Context) {
	var request dto.SignupRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if u.service == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Service is not initialized"})
		return
	}
	result, err := u.service.Register(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": result})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Success! Please wait..."})
}

// User Login
func (u *UserController) Login(c *gin.Context) {
	var request dto.LoginRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if u.service == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Service is not initialized"})
		return
	}
	result, err := u.service.Login(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Login Failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Login Success", "Data": result})
}
