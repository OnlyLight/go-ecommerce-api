package user

import "github.com/gin-gonic/gin"

type ProductRouter struct {
}

func (pr *ProductRouter) InitProductRouter(router *gin.RouterGroup) {
	productRouterPublic := router.Group("/product")
	{
		productRouterPublic.GET("/search")
		productRouterPublic.GET("/detail/:id")
	}
}
