package impl

import (
	"github.com/gin-gonic/gin"
	"server-furniture-ecommerce-gin/global"
	"server-furniture-ecommerce-gin/internal/domain/request"
	"server-furniture-ecommerce-gin/internal/repository"
	"server-furniture-ecommerce-gin/internal/service"
	"server-furniture-ecommerce-gin/pkg/exception"
	"server-furniture-ecommerce-gin/pkg/helper"
	"server-furniture-ecommerce-gin/pkg/utils/crypto"
)

type AuthServiceImpl struct {
	authRepository repository.IAuthRepository
	userRepository repository.IUserRepository
}

func NewAuthService(authRepository repository.IAuthRepository, userRepository repository.IUserRepository) service.IAuthService {
	return &AuthServiceImpl{
		authRepository: authRepository,
		userRepository: userRepository,
	}
}

func (asi *AuthServiceImpl) VerifyAccount(otp string) int {
	user := asi.authRepository.GetUserByOTP(otp)
	if user == nil {
		global.Logger.Error("Failed when search user by otp")
		return exception.NotFoundCode
	}
	user.IsActive = true
	isSave := asi.userRepository.Update(user)
	if !isSave {
		global.Logger.Error("Failed when update user")
		return exception.ErrorUpdateCode
	}
	return exception.SuccessCode
}

func (asi *AuthServiceImpl) Login(loginData *request.LoginInput) bool {
	user, _ := asi.authRepository.GetUserByUsername(loginData.Username)
	if user == nil {
		return false
	}
	hashPass := crypto.GetHash(loginData.Password)
	if hashPass == user.Password {
		return true
	}
	return false
}

func (asi *AuthServiceImpl) Logout(logoutData *request.LogoutData, ctx *gin.Context) int {
	helper.GetUserFromContext(ctx)

	panic("implement me")
}
