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

var _ interfaces.UserInterface = &UserService{}

type UserService struct {
}

func (u UserService) Register(request dto.SignupRequest) (message string, err error) {

	// Check Duplicate Email
	isDup, err := u.isDuplicateEmail(request.Email)
	if err != nil || isDup == false {
		message = "Email already exists"
		err = errors.New("Email already exists")
		return message, err
	}
	// Check Legal Password
	if !u.isLegalPassword(request.Password) {
		message = "Password must be at least 10 characters long and contain at least one uppercase letter, one lowercase letter, one number, and one special character"
		err = errors.New("Password must be at least 10 characters long and contain at least one uppercase letter, one lowercase letter, one number, and one special character")
		return message, err
	}
	isDup, err = u.isDuplicatePhoneNumber(request.PhoneNumber)
	// Check Duplicate PhoneNumber
	if err != nil || isDup == false {
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
func (u UserService) Login(request *dto.LoginRequest) (*dto.LoginResponse, error) {
	var user dto.LoginResponse
	// Check Legal Password
	if !u.isLegalPassword(request.Password) {
		return &dto.LoginResponse{}, errors.New("Password must be at least 10 characters long and contain at least one uppercase letter, one lowercase letter, one number, and one special character")
	}
	// Saltin Password with Phone Number
	newPassword := request.Password + request.Phone
	// Hash Password
	newHashedPassword := hashPassword(newPassword)
	err := db.DB.QueryRow("SELECT id,phone_number,email,full_name FROM user WHERE phone_number = ? AND password = ? AND deleted_at IS NULL", request.Phone, newHashedPassword).Scan(&user.Id, &user.PhoneNumber, &user.Email, &user.FullName)
	if err != nil {
		return &dto.LoginResponse{}, errors.New("Phone number or password is incorrect. Please try again")
	}
	return &user, nil
}
func (u UserService) TokenLogin(request *dto.LoginRequest) (string, error) {
	user, err := u.Login(request)
	if err != nil {
		return "", err
	}
	token, err := CreateToken(user)
	if err != nil {
		return "", err
	}
	return token, nil
}

// Update User Information
func (u UserService) Update(request *dto.UserUpdateRequest) (message string, err error) {
	// Check duplicate phone number
	isDup, err := isDuplicatePhoneNumberForUpdate(request.PhoneNumber, request.Id)
	if err != nil || isDup == false {
		message = "Phone number already exists"
		err = errors.New("Phone number already exists")
		return message, err
	}
	// Check duplicate email
	isDup, err = isDuplicateEmailForUpdate(request.Email, request.Id)
	if err != nil || isDup == false {
		message = "Email already exists"
		err = errors.New("Email already exists")
		return message, err
	}
	// Check Legal Password
	if !u.isLegalPassword(request.Password) {
		message = "Password must be at least 10 characters long and contain at least one uppercase letter, one lowercase letter, one number, and one special character"
		err = errors.New("Password must be at least 10 characters long and contain at least one uppercase letter, one lowercase letter, one number, and one special character")
		return message, err
	}
	newPassword := request.Password + request.PhoneNumber
	newHashedPassword := hashPassword(newPassword)
	_, err = db.DB.Exec("UPDATE user SET phone_number = ?, password =?, full_name = ?, email = ? WHERE id = ? AND deleted_at IS NULL", request.PhoneNumber, newHashedPassword, request.FullName, request.Email, request.Id)
	if err != nil {
		message = "Failed to update"
		err = errors.New("Failed to update")
		return message, err
	}
	message = "Update Success"
	return message, nil
}

// --------------------------------------------------------------------------
// Internal Code
// Check Duplicate Phone Number
func (u UserService) isDuplicatePhoneNumber(phone string) (bool, error) {
	querry, err := db.DB.Query("SELECT * FROM user WHERE phone_number = ? AND deleted_at IS NULL", phone)
	if err != nil || querry.Next() {
		return false, err
	}
	return true, nil
}

// Check Legal Password
func (u UserService) isLegalPassword(password string) bool {
	return len(password) >= 10 && regexp.MustCompile(`^[A-Za-z\d!@#$%^&*()_+{}|:<>?~]+$`).MatchString(password) &&
		regexp.MustCompile(`[a-z]`).MatchString(password) &&
		regexp.MustCompile(`[A-Z]`).MatchString(password) &&
		regexp.MustCompile(`\d`).MatchString(password) &&
		regexp.MustCompile(`[!@#$%^&*()_+{}|:<>?~]`).MatchString(password)
}

// Check Duplicate Email
func (u UserService) isDuplicateEmail(email string) (bool, error) {
	querry, err := db.DB.Query("SELECT * FROM user WHERE email =? AND deleted_at IS NULL", email)
	if err != nil || querry.Next() {
		return false, err
	}
	return true, nil
}

// Hash Password
func hashPassword(password string) string {
	passwordHash := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", passwordHash)
}

// Check Duplicate Email for User Update Information
func isDuplicateEmailForUpdate(email string, id int) (bool, error) {
	querry, err := db.DB.Query("SELECT * FROM user WHERE email =? AND id != ? AND deleted_at IS NULL", email, id)
	if err != nil || querry.Next() {
		return false, err
	}
	return true, nil
}

// Check Duplicate Phone Number for User Update Information
func isDuplicatePhoneNumberForUpdate(phone string, id int) (bool, error) {
	querry, err := db.DB.Query("SELECT * FROM user WHERE phone_number = ? AND id != ? AND deleted_at IS NULL", phone, id)
	if err != nil || querry.Next() {
		return false, err
	}
	return true, nil
}
