package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server-furniture-ecommerce-gin/global"
	"server-furniture-ecommerce-gin/internal/domain/request"
	"server-furniture-ecommerce-gin/internal/domain/response"
	"server-furniture-ecommerce-gin/internal/service"
	"server-furniture-ecommerce-gin/pkg/exception"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) Register(c *gin.Context) {
	var registerReq request.RegisterRequest
	if err := c.ShouldBind(&registerReq); err != nil {
		response.ErrorResponse(c, exception.BadRequestCode, err.Error())
		return
	}
	global.Logger.Info("Received Register Request", zap.String("email", registerReq.Email))
	result := uc.userService.Register(registerReq)
	response.SuccessResponse(c, result, nil)
}

func (uc *UserController) ChangePassword(c *gin.Context) {
	var changePasswordData request.ChangePasswordRequest
	if err := c.ShouldBind(&changePasswordData); err != nil {
		response.ErrorResponse(c, exception.BadRequestCode, err.Error())
		return
	}
	result := uc.userService.ChangePassword(changePasswordData, c)
	response.SuccessResponse(c, exception.SuccessCode, result)
}
