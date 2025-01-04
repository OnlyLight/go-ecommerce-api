package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/onlylight29/go-ecommerce-backend-api/global"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/constance"
	"github.com/onlylight29/go-ecommerce-backend-api/pkg/response"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func GetUserKey(hashKey string) string {
	return fmt.Sprintf("u:%s:otp", hashKey)
}

func HandleOTPFound(otp string, err error) (int, error) {
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

func HandleOTPValidation(ctx context.Context, hashKey, correctOTP, inputOTP string) error {
	keyAttempts := GetUserKey(hashKey)
	keyBlocked := fmt.Sprintf("u:%s:blocked", hashKey)

	// Check if the user is blocked
	blocked, err := global.RDB.Get(ctx, keyBlocked).Result()
	if err != nil && err != redis.Nil {
		return err
	}
	if blocked != "" {
		return fmt.Errorf("user is blocked until %s", blocked)
	}

	// Validate OTP
	if inputOTP == correctOTP {
		global.RDB.Del(ctx, keyAttempts) // Reset attempts on success
		fmt.Println("OTP validated successfully!")
		return nil
	}

	// Increment the attempt count
	attempts, err := global.RDB.Incr(ctx, keyAttempts).Result()
	if err != nil {
		return err
	}

	// Set expiration for attempts key
	if attempts == 1 {
		global.RDB.Expire(ctx, keyAttempts, time.Duration(constance.TIME_OTP_REGISTER)*time.Minute) // Expire after 5 minutes
	}

	// Check if attempts exceed the limit
	if attempts > 3 {
		blockTime := time.Duration(constance.TIME_OTP_BLOCK) * time.Minute // Block for 5 minutes
		global.RDB.Set(ctx, keyBlocked, time.Now().Add(blockTime).Format(time.RFC3339), blockTime)
		return fmt.Errorf("user is blocked for %s", blockTime)
	}

	return fmt.Errorf("invalid OTP. Attempt %d of 3", attempts)
}
