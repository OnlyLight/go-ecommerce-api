package account

import (
	"github.com/gin-gonic/gin"
	"github.com/onlylight29/go-ecommerce-backend-api/global"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/service"
	"github.com/onlylight29/go-ecommerce-backend-api/pkg/response"
	"go.uber.org/zap"
)

var Login = new(cUserLogin)

type cUserLogin struct {
}

func (c *cUserLogin) Login(ctx *gin.Context) {
	err := service.UserLogin().Login(ctx)
	if err != nil {
		global.Logger.Error("Can not implement UserLogin interface", zap.Error(err))
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid)
		return
	}

	response.SuccessResponse(ctx, response.ErrCodeSuccess, nil)
}
