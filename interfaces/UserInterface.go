package interfaces

import (
	"RestuarantBackend/models"
	dto "RestuarantBackend/models/dto"
)

type UserInterface interface {
	Login(loginRequest *dto.LoginRequest) (*dto.LoginResponse, error)
	Register(RegisterRequest dto.SignupRequest) (string, error)
	Update(updateRequest *dto.UserUpdateRequest) (string, error)
	TokenLogin(loginRequest *dto.LoginRequest) (string, error)
	PagingListAllUser(pagingRequest *dto.PagingRequest) ([]models.User, error)
	GetAllUser() ([]models.User, error)
	BlockOrUnBlockUser(userId *int) (string, error)
}
