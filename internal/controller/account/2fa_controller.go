package account

import (
	"github.com/gin-gonic/gin"
	"github.com/onlylight29/go-ecommerce-backend-api/global"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/model"
	"github.com/onlylight29/go-ecommerce-backend-api/pkg/response"
	"go.uber.org/zap"
)

var TwoFA = new(sUser2FA)

type sUser2FA struct{}

func (s *sUser2FA) SetupTwoFactorAuth(ctx *gin.Context) {
	var params model.SetupTwoFactorAuthInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		global.Logger.Error("Can Bind JSON", zap.Error(err))
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthSetupFailed)
		return
	}
}
