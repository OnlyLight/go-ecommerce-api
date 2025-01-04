package user

import (
	"github.com/gin-gonic/gin"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/controller/account"
)

type UserRouter struct {
}

func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	// userController, _ := wire.InitUserRouterHandler()

	// Public routes
	userRouterPublic := router.Group("/user")
	{
		// userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.POST("/register", account.Login.Register)
		userRouterPublic.POST("/verify_account", account.Login.VerifyOTP)
		userRouterPublic.POST("/login", account.Login.Login)
	}

	// Private routes
	userRouterPrivate := router.Group("/user")
	{
		userRouterPrivate.GET("/get_info")
	}
}
