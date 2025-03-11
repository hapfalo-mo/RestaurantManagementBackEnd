package service

import (
	dto "RestuarantBackend/models/dto"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var sercretKey = []byte("ChuonNG_deP_trAi")

// Create Struct Claims
type Claims struct {
	UserID   int    `json:"userID`
	Role     int    `json:"role"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	FullName string `json:"fullName"`
	Point    int    `json:"point"`
	jwt.RegisteredClaims
}

// Create Token
func CreateToken(user *dto.LoginResponse) (string, error) {
	claims := &Claims{
		UserID:   user.Id,
		Email:    user.Email,
		Phone:    user.PhoneNumber,
		FullName: user.FullName,
		Role:     user.Role,
		Point:    user.Point,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 1)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "RestuarantBackend",
			Subject:   "Authentication",
		},
	}
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(sercretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return sercretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if errors.Is(err, jwt.ErrTokenExpired) {
		return nil, errors.New("Token Expired")
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, err
	}
	return claims, nil
}
