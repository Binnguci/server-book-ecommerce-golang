package impl

import (
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"server-furniture-ecommerce-gin/global"
	"server-furniture-ecommerce-gin/internal/domain/request"
	"server-furniture-ecommerce-gin/internal/model"
	"server-furniture-ecommerce-gin/internal/repository"
	"server-furniture-ecommerce-gin/internal/service"
	"server-furniture-ecommerce-gin/pkg/exception"
	"server-furniture-ecommerce-gin/pkg/utils/crypto"
	"server-furniture-ecommerce-gin/pkg/utils/random"
	"server-furniture-ecommerce-gin/pkg/utils/send"
	"strconv"
	"time"
)

type UserServiceImpl struct {
	userRepository repository.IUserRepository
	roleRepository repository.IRoleRepository
	authRepository repository.IAuthRepository
}

func NewUserService(userRepo repository.IUserRepository, roleRepo repository.IRoleRepository, authRepo repository.IAuthRepository) service.IUserService {
	return &UserServiceImpl{
		userRepository: userRepo,
		roleRepository: roleRepo,
		authRepository: authRepo,
	}
}

func (usi *UserServiceImpl) Register(register request.RegisterRequest) int {
	if usi.userRepository.GetUserByEmail(register.Email) {
		return exception.UserExistsCode
	}

	user := usi.mapRequestToUser(register)

	role, err := usi.getDefaultRole()
	if err != nil {
		return exception.NotFoundCode
	}
	user.RoleID = role.ID

	otp := usi.generateAndSetOTP(user)

	otpInt, err := strconv.Atoi(otp)
	if err != nil {
		return exception.ErrorInvalidOTP
	}

	expirationTime := time.Now().Add(3 * time.Minute).Unix()

	err = usi.authRepository.AddOTP(user.Email, otpInt, expirationTime)
	if err != nil {
		return exception.ErrorInvalidOTP
	}
	hashedPassword := crypto.GetHash(user.Password)
	user.Password = hashedPassword

	if err := usi.sendOTPEmail(register.Email, otp); err != nil {
		return exception.ErrorSendEmail
	}

	if !usi.userRepository.Register(user) {
		return exception.CreateFailedCode
	}
	return exception.CreateSuccessCode
}

func (usi *UserServiceImpl) mapRequestToUser(register request.RegisterRequest) *model.User {
	user := &model.User{}
	copier.Copy(user, &register)
	return user
}

func (usi *UserServiceImpl) getDefaultRole() (*model.Role, error) {
	role, err := usi.roleRepository.GetRoleByName("USER")
	if err != nil {
		global.Logger.Error("Role not found", zap.Error(err))
		return nil, err
	}
	return role, nil
}

func (usi *UserServiceImpl) generateAndSetOTP(user *model.User) string {
	otp := random.GenerateOTP()
	user.Otp = strconv.Itoa(otp)
	user.OtpExpired = time.Now().Add(3 * time.Minute)
	global.Logger.Info("Generated OTP", zap.String("otp", user.Otp))
	return user.Otp
}

func (usi *UserServiceImpl) sendOTPEmail(email string, otp string) error {
	mail := global.Config.Mail
	if err := send.SendOTPToEmail([]string{email}, mail.Username, otp); err != nil {
		global.Logger.Error("Failed to send OTP email", zap.Error(err))
		return err
	}
	return nil
}
