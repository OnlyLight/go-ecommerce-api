package user

import "github.com/gin-gonic/gin"

type UserRouter struct {
}

func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	// Public routes
	userRouterPublic := router.Group("/user")
	{
		userRouterPublic.POST("/register")
		userRouterPublic.POST("/otp")
	}

	// Private routes
	userRouterPrivate := router.Group("/user")
	{
		userRouterPrivate.GET("/get_info")
	}
}
