package service

import (
	"server-furniture-ecommerce-gin/internal/domain/request"
)

type IUserService interface {
	Register(register request.RegisterRequest) int
}
