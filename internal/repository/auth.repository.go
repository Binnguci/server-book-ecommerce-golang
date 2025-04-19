package repository

import (
	"server-furniture-ecommerce-gin/internal/model"
)

type IAuthRepository interface {
	AddOTP(email string, otp int, expirationTime int64) error
	GetUserByOTP(otp string) (*model.User, error)
	GetUserByUsernameAndPassword(username string, password string) (bool, error)
}
