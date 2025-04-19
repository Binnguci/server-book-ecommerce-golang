package impl

import (
	"go.uber.org/zap"
	"server-furniture-ecommerce-gin/global"
	"server-furniture-ecommerce-gin/internal/domain/request"
	"server-furniture-ecommerce-gin/internal/repository"
	"server-furniture-ecommerce-gin/internal/service"
	"server-furniture-ecommerce-gin/pkg/exception"
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
	user, err := asi.authRepository.GetUserByOTP(otp)
	if err != nil {
		global.Logger.Error("Failed when search user by otp", zap.Error(err))
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

func (asi *AuthServiceImpl) Login(loginData *request.LoginInput) (bool, error) {
	_, err := asi.authRepository.GetUserByUsernameAndPassword(loginData.Username, loginData.Password)
	if err != nil {
		return false, err
	}
	return true, nil
}
