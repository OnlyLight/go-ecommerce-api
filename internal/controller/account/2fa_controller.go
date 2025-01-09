package account

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/onlylight29/go-ecommerce-backend-api/global"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/model"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/service"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/utils/context"
	"github.com/onlylight29/go-ecommerce-backend-api/pkg/response"
	"go.uber.org/zap"
)

var TwoFA = new(sUser2FA)

type sUser2FA struct{}

// User SetupTwoFactorAuth doc
// @Summary      User SetupTwoFactorAuth
// @Description  SetupTwoFactorAuth
// @Tags         account 2fa
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Authorization token"
// @Param        payload body model.SetupTwoFactorAuthInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Router       /user/two-factor/setup [post]
func (s *sUser2FA) SetupTwoFactorAuth(ctx *gin.Context) {
	var params model.SetupTwoFactorAuthInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		global.Logger.Error("Can Bind JSON", zap.Error(err))
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthSetupFailed)
		return
	}

	userId, err := context.GetUserIdFromUUID(ctx.Request.Context())
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthSetupFailed)
		return
	}

	log.Println("UserId::SetupTwoFactorAuth::", userId)
	params.UserId = uint32(userId)
	codeResult, err := service.UserLogin().SetupTwoFactorAuth(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthSetupFailed)
	}

	response.SuccessResponse(ctx, codeResult, nil)
}

// User VerifyTwoFactorAuth doc
// @Summary      User VerifyTwoFactorAuth
// @Description  VerifyTwoFactorAuth
// @Tags         account 2fa
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Authorization token"
// @Param        payload body model.TwoFactorVerificationInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Router       /user/two-factor/verify [post]
func (s *sUser2FA) VerifyTwoFactorAuth(ctx *gin.Context) {
	var params model.TwoFactorVerificationInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		global.Logger.Error("Can Bind JSON", zap.Error(err))
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthVerifyFailed)
		return
	}

	userId, err := context.GetUserIdFromUUID(ctx.Request.Context())
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthVerifyFailed)
		return
	}

	log.Println("UserId::VerifyTwoFactorAuth::", userId)
	params.UserId = uint32(userId)
	codeResult, err := service.UserLogin().VerifyTwoFactorAuth(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthVerifyFailed)
	}

	response.SuccessResponse(ctx, codeResult, nil)
}
