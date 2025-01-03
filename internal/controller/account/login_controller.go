package account

import (
	"github.com/gin-gonic/gin"
	"github.com/onlylight29/go-ecommerce-backend-api/global"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/model"
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

// User Registeration doc
// @Summary      User Registeration
// @Description  When user is registered send otp to email
// @Tags         account management
// @Accept       json
// @Produce      json
// @Param        payload body model.RegisterInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Router       /user/register [post]
func (c *cUserLogin) Register(ctx *gin.Context) {
	var params model.RegisterInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		global.Logger.Error("Can Bind JSON", zap.Error(err))
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid)
		return
	}

	codeStatus, err := service.UserLogin().Register(ctx, &params)
	if err != nil {
		global.Logger.Error("Can not implement UserLogin interface", zap.Error(err))
		response.ErrorResponse(ctx, codeStatus)
		return
	}

	response.SuccessResponse(ctx, codeStatus, nil)
}
