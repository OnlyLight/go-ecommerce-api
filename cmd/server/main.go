package main

import (
	"fmt"

	_ "github.com/onlylight29/go-ecommerce-backend-api/docs"
	"github.com/onlylight29/go-ecommerce-backend-api/global"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/initialize"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           API Documentation Ecommerce Backend
// @version         1.0.0
// @description     This is a sample server celler server.
// @termsOfService  https://github.com/OnlyLight/go-ecommerce-api

// @contact.name   API Support
// @contact.url    github.com/OnlyLight/go-ecommerce-api
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8002
// @BasePath  /v1/api
func main() {
	r := initialize.Run()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(fmt.Sprintf(":%d", global.Config.Server.Port))
}
