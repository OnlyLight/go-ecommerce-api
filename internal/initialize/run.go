package initialize

import (
	"fmt"

	"github.com/onlylight29/go-ecommerce-backend-api/global"
)

func Run() {
	LoadConfig()
	// fmt.Println("config", global.Config.Server.Port)

	InitLogger()
	global.Logger.Info("Logger init success")

	InitMySQL()
	InitRedis()

	r := InitRouter()

	r.Run(fmt.Sprintf(":%d", global.Config.Server.Port))
}
