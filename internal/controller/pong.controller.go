package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/service"
)

type PongController struct {
	pongService *service.PongService
}

func NewPongController() *PongController {
	return &PongController{
		pongService: service.NewPongService(),
	}
}

func (pc *PongController) GetPongInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": pc.pongService.GetInfoPongService(),
	})
}
