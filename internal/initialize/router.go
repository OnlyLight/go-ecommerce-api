package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/controller"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/middlewares"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.AuthenMiddleware())
	r.GET("/ping", controller.NewPongController().GetPongInfo)
	r.GET("/user/:uid", controller.NewUserController().GetUserByID)

	return r
}
