package repository

import (
	"server-furniture-ecommerce-gin/internal/model"
)

type IAuthRepository interface {
	AddOTP(email string, otp int, expirationTime int64) error
	GetUserByOTP(otp string) *model.User
	GetUserByUsernameAndPassword(username string, password string) bool
	GetUserByUsername(username string) (*model.User, error)
	CheckPassword(password string) bool
	SaveTokenInvalid(tokenInvalid *string) bool
}
