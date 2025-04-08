package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/onlylight29/go-ecommerce-backend-api/global"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/middlewares"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/routers"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default() // Show log debug on console
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// Middleware
	// r.use() // logger
	// r.use() // cors
	// r.use() // limit request global
	manageRouter := routers.RouterGroupApp.Manage
	userRouter := routers.RouterGroupApp.User

	r.Use(middlewares.NewRateLimiter().GlobalRateLimiter())
	r.Use(middlewares.NewRateLimiter().PublicAPIRateLimiter())
	r.Use(middlewares.NewRateLimiter().UserPrivateAPIRateLimiter())

	MainGroup := r.Group("/v1/api")
	{
		MainGroup.GET("/checkStatus")
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
		userRouter.InitTicketRouter(MainGroup)
	}
	{
		manageRouter.InitUserRouter(MainGroup)
		manageRouter.InitAdminRouter(MainGroup)
	}

	return r
}
