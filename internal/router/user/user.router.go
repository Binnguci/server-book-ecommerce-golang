package user

import (
	"github.com/gin-gonic/gin"
	"server-furniture-ecommerce-gin/internal/injector"
	"server-furniture-ecommerce-gin/internal/middleware"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	userController, _ := injector.InitUserControllerHandler()
	authController, _ := injector.InitAuthControllerHandler()

	//public router
	userRouterPublic := Router.Group("/auth")
	{
		userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.GET("/confirm-account", authController.VerifyAccount)
		userRouterPublic.POST("/login", authController.Login)
	}

	//private router
	userRouterPrivate := Router.Group("/auth")
	userRouterPrivate.Use(middleware.AuthenticationMiddleware())
	userRouterPrivate.Use(middleware.PermissionMiddleware())
	userRouterPrivate.Use(middleware.JWTAuthMiddleware())
	{

	}
}
