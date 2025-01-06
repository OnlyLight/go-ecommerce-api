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

type cUserLogin struct{}

// User Login doc
// @Summary      User Login
// @Description  Login
// @Tags         account management
// @Accept       json
// @Produce      json
// @Param        payload body model.LoginInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Router       /user/login [post]
func (c *cUserLogin) Login(ctx *gin.Context) {
	var params model.LoginInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		global.Logger.Error("Can Bind JSON", zap.Error(err))
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid)
		return
	}

	codeResult, dataResult, err := service.UserLogin().Login(ctx, &params)
	if err != nil {
		global.Logger.Error("Can not implement UserLogin interface", zap.Error(err))
		response.ErrorResponse(ctx, codeResult)
		return
	}

	response.SuccessResponse(ctx, codeResult, dataResult)
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
		global.Logger.Error("Can not implement UserLogin.Register interface", zap.Error(err))
		response.ErrorResponse(ctx, codeStatus)
		return
	}

	response.SuccessResponse(ctx, codeStatus, nil)
}

// User VerifyOTP doc
// @Summary      User VerifyOTP
// @Description  Verify OTP
// @Tags         account management
// @Accept       json
// @Produce      json
// @Param        payload body model.VerifyInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Router       /user/verify_account [post]
func (c *cUserLogin) VerifyOTP(ctx *gin.Context) {
	var params model.VerifyInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		global.Logger.Error("Can Bind JSON", zap.Error(err))
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid)
		return
	}

	out, err := service.UserLogin().VerifyOTP(ctx, &params)
	if err != nil {
		global.Logger.Error("Can not implement UserLogin.VerifyOTP interface", zap.Error(err))
		response.ErrorResponse(ctx, response.ErrInvalOTP)
		return
	}

	response.SuccessResponse(ctx, response.ErrCodeSuccess, out)
}

// User UpdatePasswordRegister doc
// @Summary      User UpdatePasswordRegister
// @Description  Update Password Register
// @Tags         account management
// @Accept       json
// @Produce      json
// @Param        payload body model.UpdatePasswordRegisterInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Router       /user/update_pass_register [post]
func (c *cUserLogin) UpdatePasswordRegister(ctx *gin.Context) {
	var params model.UpdatePasswordRegisterInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		global.Logger.Error("Can Bind JSON", zap.Error(err))
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid)
		return
	}

	userid, err := service.UserLogin().UpdatePasswordRegister(ctx, params.UserToken, params.UserPassword)
	if err != nil {
		global.Logger.Error("Can not implement UserLogin.UpdatePasswordRegister interface", zap.Error(err))
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid)
		return
	}

	response.SuccessResponse(ctx, response.ErrCodeSuccess, userid)
}
