package service

import (
	"github.com/gin-gonic/gin"
	"server-furniture-ecommerce-gin/internal/domain/request"
)

type IAuthService interface {
	VerifyAccount(otp string) int
	Login(loginData *request.LoginInput) bool
	Logout(logoutData *request.LogoutData, c *gin.Context) int
}
