package user

import "github.com/gin-gonic/gin"

type ProductRouter struct {
}

func (pr *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	productRouterPublic := Router.Group("/product")
	{
		productRouterPublic.GET("/search/:url")
	}
}
