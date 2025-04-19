package repository

import (
	"server-furniture-ecommerce-gin/internal/model"
)

type IUserRepository interface {
	GetUserByEmail(email string) bool
	Register(user *model.User) bool
	Update(user *model.User) bool
}
