package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/onlylight29/go-ecommerce-backend-api/global"
)

func Run() *gin.Engine {
	LoadConfig()
	// fmt.Println("config", global.Config.Server.Port)

	InitLogger()
	global.Logger.Info("Logger init success")

	// InitMySQL()
	InitMySQLC()
	InitServiceInterface()

	InitRedis()
	InitKafka()

	return InitRouter()
}
