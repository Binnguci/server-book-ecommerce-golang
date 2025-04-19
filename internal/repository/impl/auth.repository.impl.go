package impl

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"server-furniture-ecommerce-gin/global"
	"server-furniture-ecommerce-gin/internal/model"
	"server-furniture-ecommerce-gin/internal/repository"
	"server-furniture-ecommerce-gin/pkg/utils/crypto"
	"time"
)

type AuthRepositoryImpl struct {
}

func NewAuthRepository() repository.IAuthRepository {
	return &AuthRepositoryImpl{}
}

func (ari *AuthRepositoryImpl) AddOTP(email string, otp int, expirationTime int64) error {
	key := fmt.Sprint("user:%s:otp", email)
	return global.Rdb.SetEx(ctx, key, otp, time.Duration(expirationTime)).Err()
}

func (ari *AuthRepositoryImpl) GetUserByOTP(otp string) *model.User {
	user := &model.User{}
	err := global.Mdb.Table(model.TableNameUser).Where("otp = ?", otp).First(user).Error
	if err != nil {
		return nil
	}
	return user
}

func (ari *AuthRepositoryImpl) GetUserByUsernameAndPassword(username string, password string) bool {
	var user model.User
	err := global.Mdb.Table(model.TableNameUser).Where("username = ?", username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	if err != nil {
		return false
	}
	hashedInputPassword := crypto.GetHash(password)
	if user.Password != hashedInputPassword {
		return false
	}
	return true
}

func (ari *AuthRepositoryImpl) GetUserByUsername(username string) (*model.User, error) {
	user := &model.User{}
	err := global.Mdb.Table(model.TableNameUser).Where("username  = ?", username).First(user)
	if err != nil {
		return nil, err.Error
	}
	return user, nil
}

func (ari *AuthRepositoryImpl) CheckPassword(password string) bool {
	//TODO implement me
	panic("implement me")
}

func (ari *AuthRepositoryImpl) SaveTokenInvalid(tokenInvalid *string) bool {
	err := global.Mdb.Table(model.TableNameInvalidatedToken).Create(&tokenInvalid)
	if err != nil {
		return false
	}
	return true
}
