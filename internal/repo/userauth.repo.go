package repo

import (
	"fmt"
	"time"

	"github.com/onlylight29/go-ecommerce-backend-api/global"
)

type IUserAuthRepository interface {
	AddOTP(email string, otp int, expirationTime int64) error
}

type userAuthRepository struct{}

func NewUserAuthRepository() IUserAuthRepository {
	return &userAuthRepository{}
}

// AddOTP implements IUserAuthRepository.
func (u *userAuthRepository) AddOTP(email string, otp int, expirationTime int64) error {
	key := fmt.Sprintf("usr:%s!otp", email)
	return global.RDB.SetEx(ctx, key, otp, time.Duration(expirationTime)).Err()
}
