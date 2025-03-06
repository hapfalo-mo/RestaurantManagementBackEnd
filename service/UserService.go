package service

import (
	"RestuarantBackend/db"
	"RestuarantBackend/interfaces"
	dto "RestuarantBackend/models/dto"
	"crypto/sha256"
	"errors"
	"fmt"
	"regexp"
)

var _interfaces interfaces.UserInterface = &UserService{}

type UserService struct {
}

func (u UserService) Register(request dto.SignupRequest) (message string, err error) {

	// Check Duplicate Email
	if !u.CheckDuplicateEmail(request.Email) {
		message = "Email already exists"
		err = errors.New("Email already exists")
		return message, err
	}
	// Check Legal Password
	if !u.CheckLegalPassword(request.Password) {
		message = "Password must be at least 10 characters long and contain at least one uppercase letter, one lowercase letter, one number, and one special character"
		err = errors.New("Password must be at least 10 characters long and contain at least one uppercase letter, one lowercase letter, one number, and one special character")
		return message, err
	}

	// Check Duplicate PhoneNumber
	if !u.CheckDuplicatePhoneNumber(request.PhoneNumber) {
		message = "Phone number already exists"
		err = errors.New("Phone number already exists")
		return message, err
	}
	// Salting Password
	newPassword := request.Password + request.PhoneNumber
	// Hash Password
	newHashedPassword := hashPassword(newPassword)

	_, err = db.DB.Exec("INSERT INTO `user` (phone_number, password, email, full_name) VALUES (?,?,?,?)", request.PhoneNumber, newHashedPassword, request.Email, request.FullName)
	if err != nil {
		message = "Failed to register"
		err = errors.New("Failed to register")
		return message, err
	}
	message = "Register Success"
	return message, nil
}

// Login Function for User
func (u UserService) Login(request dto.LoginRequest) (dto.LoginResponse, error) {
	var user dto.LoginResponse
	isDup, err := u.CheckDuplicatePhoneNumber(request.Phone)
	if err != nil {
		return dto.LoginResponse{}, err
	}
	// Check Duplicate Phone Number
	if isDup {
		return dto.LoginResponse{}, errors.New("Phone number not found")
	}
	// Check Legal Password
	if !u.CheckLegalPassword(request.Password) {
		return dto.LoginResponse{}, errors.New("Password must be at least 10 characters long and contain at least one uppercase letter, one lowercase letter, one number, and one special character")
	}
	// Saltin Password with Phone Number
	newPassword := request.Password + request.Phone
	// Hash Password
	newHashedPassword := hashPassword(newPassword)
	err = db.DB.QueryRow("SELECT id,phone_number,email,full_name FROM user WHERE phone_number = ? AND password = ?", request.Phone, newHashedPassword).Scan(&user.Id, &user.PhoneNumber, &user.Email, &user.FullName)
	if err != nil {
		return dto.LoginResponse{}, errors.New("Phone number or password is incorrect. Please try again")
	}
	return user, nil
}

// Check Duplicate Phone Number
func (u UserService) CheckDuplicatePhoneNumber(phone string) (bool, error) {
	querry, err := db.DB.Query("SELECT * FROM user WHERE phone_number = ? AND deleted_at IS NULL", phone)
	if err != nil {
		return false, err
	}
	if querry.Next() {
		return false, nil
	}
	return true, nil
}

// Check Legal Password
func (u UserService) CheckLegalPassword(password string) bool {
	return len(password) < 10 && regexp.MustCompile(`(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*()_+{}|:<>?~])`).MatchString(password)
}

// Check Duplicate Email
func (u UserService) CheckDuplicateEmail(email string) bool {
	querry, err := db.DB.Query("SELECT * FROM user WHERE email =? AND deleted_at IS NULL", email)
	if err != nil {
		return false
	}
	if querry.Next() {
		return false
	}
	return true
}

// Internal Code
// Hash Password
func hashPassword(password string) string {
	passwordHash := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", passwordHash)
}
