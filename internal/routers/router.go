package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/controller"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", controller.NewPongController().GetPongInfo)
	r.GET("/user/:uid", controller.NewUserController().GetUserByID)

	return r
}
