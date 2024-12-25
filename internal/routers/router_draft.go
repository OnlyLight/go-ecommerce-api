package routers

// import (
// 	"fmt"

// 	"github.com/gin-gonic/gin"
// 	"github.com/onlylight29/go-ecommerce-backend-api/internal/controller"
// 	"github.com/onlylight29/go-ecommerce-backend-api/internal/middlewares"
// )

// func AA() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		fmt.Println("Before --> AA")
// 		ctx.Next()
// 		fmt.Println("After --> AA")
// 	}
// }
// func BB() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		fmt.Println("Before --> BB")
// 		ctx.Next()
// 		fmt.Println("After --> BB")
// 	}
// }
// func CC(ctx *gin.Context) {
// 	fmt.Println("Before --> CC")
// 	ctx.Next()
// 	fmt.Println("After --> CC")
// }

// func NewRouter() *gin.Engine {
// 	r := gin.Default()
// 	r.Use(middlewares.AuthenMiddleware(), AA(), BB(), CC)
// 	r.GET("/ping", controller.NewPongController().GetPongInfo)
// 	r.GET("/user/:uid", controller.NewUserController().GetUserByID)

// 	return r
// }
