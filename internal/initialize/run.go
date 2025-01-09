package initialize

import (
	"github.com/gin-gonic/gin"
)

func Run() *gin.Engine {
	LoadConfig()
	// fmt.Println("config", global.Config.Server.Port)

	InitLogger()

	// InitMySQL()
	InitMySQLC()
	InitServiceInterface()

	InitRedis()
	InitKafka()

	return InitRouter()
}
