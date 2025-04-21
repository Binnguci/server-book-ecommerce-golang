package controller

import (
	"github.com/gin-gonic/gin"
	"server-furniture-ecommerce-gin/internal/constant"
	"server-furniture-ecommerce-gin/internal/domain/request"
	"server-furniture-ecommerce-gin/internal/domain/response"
	"server-furniture-ecommerce-gin/internal/service"
	"server-furniture-ecommerce-gin/pkg/exception"
	"server-furniture-ecommerce-gin/pkg/utils/auth"
)

type AuthController struct {
	authService service.IAuthService
}

func NewAuthController(authService service.IAuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (ac *AuthController) VerifyAccount(c *gin.Context) {
	otp := c.Query("otp")
	if err := c.ShouldBind(otp); err != nil {
		response.ErrorResponse(c, exception.BadRequestCode, err.Error())
		return
	}
	result := ac.authService.VerifyAccount(otp)
	response.SuccessResponse(c, exception.SuccessCode, result)
}

func (ac *AuthController) Login(c *gin.Context) {
	var loginData request.LoginInput

	if err := c.ShouldBindJSON(&loginData); err != nil {
		response.ErrorResponse(c, exception.BadRequestCode, constant.DATA_INVALIDATED)
		return
	}
	ok := ac.authService.Login(&loginData)
	if !ok {
		response.ErrorResponse(c, exception.NotFoundCode, "")
	}
	token, err := auth.GenerateToken(loginData.Username)
	if err != nil {
		response.ErrorResponse(c, exception.InternalServerErrorCode, "")
		return
	}

	response.SuccessResponse(c, exception.SuccessCode, response.LoginResponse{token, constant.LOGIN_SUCCESSFUL})
}
