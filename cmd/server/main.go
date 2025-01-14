package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/onlylight29/go-ecommerce-backend-api/docs"
	"github.com/onlylight29/go-ecommerce-backend-api/global"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/initialize"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_request_count_total",
		Help: "Total number of ping requests",
	},
)

func ping(c *gin.Context) {
	pingCounter.Inc()
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

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

	prometheus.MustRegister(pingCounter)

	r.GET("/ping/200", ping)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.Run(fmt.Sprintf(":%d", global.Config.Server.Port))
}
