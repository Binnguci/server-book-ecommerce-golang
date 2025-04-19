package service

import (
	"server-furniture-ecommerce-gin/internal/domain/request"
)

type IAuthService interface {
	VerifyAccount(otp string) int
	Login(loginData *request.LoginInput) (bool, error)
}
