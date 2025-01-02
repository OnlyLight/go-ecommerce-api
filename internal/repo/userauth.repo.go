package repo

import (
	"time"

	"github.com/onlylight29/go-ecommerce-backend-api/global"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/utils"
)

type IUserAuthRepository interface {
	AddOTP(email string, otp int, expirationTime int64) error
}

type userAuthRepository struct{}

func NewUserAuthRepository() IUserAuthRepository {
	return &userAuthRepository{}
}

// AddOTP implements IUserAuthRepository.
func (u *userAuthRepository) AddOTP(hashEmail string, otp int, expirationTime int64) error {
	key := utils.GetUserKey(hashEmail)
	return global.RDB.SetEx(ctx, key, otp, time.Duration(expirationTime)).Err()
}
