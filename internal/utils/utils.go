package utils

import (
	"fmt"

	"github.com/onlylight29/go-ecommerce-backend-api/global"
	"github.com/onlylight29/go-ecommerce-backend-api/pkg/response"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func GetUserKey(hashKey string) string {
	return fmt.Sprintf("u:%s:otp", hashKey)
}

func GetOtpFromRedis(otp string, err error) (int, error) {
	if err != nil {
		if err != redis.Nil {
			global.Logger.Error("Get OTP failed", zap.Error(err))
			return response.ErrInvalOTP, err
		}
		global.Logger.Info("OTP not exist")
		return response.ErrCodeSuccess, nil
	}

	if otp != "" {
		global.Logger.Error("OTP exists but not registered")
		return response.ErrCodeOtpAlreadyExists, fmt.Errorf("OTP exists but not registered")
	}

	return response.ErrCodeSuccess, nil
}
