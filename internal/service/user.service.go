package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/onlylight29/go-ecommerce-backend-api/internal/repo"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/utils/crypto"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/utils/random"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/utils/sendto"
	"github.com/onlylight29/go-ecommerce-backend-api/pkg/response"
)

// type UserService struct {
// 	userRepo *repo.UserRepo
// }

// func NewUserService() *UserService {
// 	return &UserService{
// 		userRepo: repo.NewUserRepo(),
// 	}
// }

// func (us *UserService) GetInfoUserService() []string {
// 	return us.userRepo.GetInfoUserRepo()
// }

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo     repo.IUserRepo
	userAuthRepo repo.IUserAuthRepository
}

func NewUserService(userRepo repo.IUserRepo, userAuthRepo repo.IUserAuthRepository) IUserService {
	return &userService{
		userRepo:     userRepo,
		userAuthRepo: userAuthRepo,
	}
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	// 0. Hash email.
	hashEmail := crypto.GetHash(email)
	fmt.Println("Hash email::", hashEmail)

	// 5. check OTP available

	// 6. user spam

	// 1. Check if email is exist.
	if us.userRepo.GetUserByEmailSQLC(email) {
		return response.ErrCodeUserAlreadyExists
	}

	// 2. Generate OTP.
	otp := random.GenerateSixDigitOTP()
	if purpose == "TEST_USER" {
		otp = 123456
	}

	fmt.Println("OTP::", otp)

	// 3. Save OTP to Redis.
	err := us.userAuthRepo.AddOTP(email, otp, int64(10*time.Minute))

	if err != nil {
		return response.ErrInvalOTP
	}

	// 4. Send OTP to email.
	// err = sendto.SendTextEmailOtp([]string{email}, "onlylight@gmail.com", strconv.Itoa(otp))
	err = sendto.SendTemplateEmailOtp([]string{email}, "onlylight@gmail.com", "otp-authen.html", map[string]interface{}{
		"otp": strconv.Itoa(otp),
	})
	if err != nil {
		return response.ErrSendEmailOtp
	}

	return response.ErrCodeParamInvalid
}
