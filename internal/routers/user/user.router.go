package user

import (
	"github.com/gin-gonic/gin"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/controller/account"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/middlewares"
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
		userRouterPublic.POST("/update_pass_register", account.Login.UpdatePasswordRegister)
		userRouterPublic.POST("/login", account.Login.Login)
	}

	// Private routes
	userRouterPrivate := router.Group("/user")
	userRouterPrivate.Use(middlewares.AuthenMiddleware())
	{
		userRouterPrivate.GET("/get_info")
		userRouterPrivate.POST("/two-factor/setup", account.TwoFA.SetupTwoFactorAuth)
		userRouterPrivate.POST("/two-factor/verify", account.TwoFA.VerifyTwoFactorAuth)
	}
}
