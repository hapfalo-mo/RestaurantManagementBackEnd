package interfaces

import dto "RestuarantBackend/models/dto"

type UserInterface interface {
	Login(loginRequest *dto.LoginRequest) (*dto.LoginResponse, error)
	Register(RegisterRequest dto.SignupRequest) (string, error)
}
