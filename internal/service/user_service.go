package service

import (
	"context"

	"github.com/onlylight29/go-ecommerce-backend-api/internal/model"
)

// Just Leader can modify this file
type (
	IUserLogin interface {
		Login(ctx context.Context, in *model.LoginInput) (codeResult int, out model.LoginOutput, err error)
		Register(ctx context.Context, in *model.RegisterInput) (codeResult int, err error)
		VerifyOTP(ctx context.Context, in *model.VerifyInput) (out model.VerifyOTPOutput, err error)
		UpdatePasswordRegister(ctx context.Context, token string, password string) (userid int, err error)

		IsTwoFactorEnabled(ctx context.Context, userId int) (codeResult int, rs bool, err error)
		SetupTwoFactorAuth(ctx context.Context, in *model.SetupTwoFactorAuthInput) (codeResult int, err error)
		VerifyTwoFactorAuth(ctx context.Context, in *model.TwoFactorVerificationInput) (codeResult int, err error)
	}

	IUserInfo interface {
		GetInfoByUserId(ctx context.Context) error
		GetAllUser(ctx context.Context) error
	}

	IUserAdmin interface {
		RemoveUser(ctx context.Context) error
		FindOneUser(ctx context.Context) error
	}
)

var (
	localUserAdmin IUserAdmin
	localUserInfo  IUserInfo
	localUserLogin IUserLogin
)

// UserAdmin
func UserAdmin() IUserAdmin {
	if localUserAdmin == nil {
		panic("implement localUserAdmin not found from Interface IUserAdmin")
	}

	return localUserAdmin
}

func InitUserAdmin(i IUserAdmin) {
	localUserAdmin = i
}

// UserInfo
func UserInfo() IUserInfo {
	if localUserInfo == nil {
		panic("implement localUserInfo not found from Interface IUserInfo")
	}

	return localUserInfo
}

func InitUserInfo(i IUserInfo) {
	localUserInfo = i
}

// UserLogin
func UserLogin() IUserLogin {
	if localUserLogin == nil {
		panic("implement localUserLogin not found from Interface IUserLogin")
	}

	return localUserLogin
}

func InitUserLogin(i IUserLogin) {
	localUserLogin = i
}
