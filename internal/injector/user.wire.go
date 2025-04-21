//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"server-furniture-ecommerce-gin/internal/controller"
	repository "server-furniture-ecommerce-gin/internal/repository/impl"
	service "server-furniture-ecommerce-gin/internal/service/impl"
)

func InitUserControllerHandler() (*controller.UserController, error) {
	wire.Build(
		repository.NewUserRepository,
		repository.NewRoleRepository,
		repository.NewAuthRepository,
		service.NewUserService,
		controller.NewUserController,
	)
	return &controller.UserController{}, nil
}
