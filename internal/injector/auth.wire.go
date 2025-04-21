//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"server-furniture-ecommerce-gin/internal/controller"
	repository "server-furniture-ecommerce-gin/internal/repository/impl"
	service "server-furniture-ecommerce-gin/internal/service/impl"
)

func InitAuthControllerHandler() (*controller.AuthController, error) {
	wire.Build(
		repository.NewAuthRepository,
		repository.NewUserRepository,
		service.NewAuthService,
		controller.NewAuthController,
	)
	return &controller.AuthController{}, nil
}
