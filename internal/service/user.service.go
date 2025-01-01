package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/onlylight29/go-ecommerce-backend-api/global"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/repo"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/utils/crypto"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/utils/random"
	"github.com/onlylight29/go-ecommerce-backend-api/pkg/response"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

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

	// 4.1 Send OTP to email.
	// err = sendto.SendTextEmailOtp([]string{email}, "onlylight@gmail.com", strconv.Itoa(otp))
	// err = sendto.SendTemplateEmailOtp([]string{email}, "onlylight@gmail.com", "otp-authen.html", map[string]interface{}{
	// 	"otp": strconv.Itoa(otp),
	// })
	// if err != nil {
	// 	return response.ErrSendEmailOtp
	// }

	// 4.2 Send OTP to Kafka
	body := make(map[string]interface{})
	body["otp"] = otp
	body["email"] = email

	jsonBody, _ := json.Marshal(body)

	message := kafka.Message{
		Key:   []byte("otp-auth"),
		Value: []byte(jsonBody),
		Time:  time.Now(),
	}

	err = global.KafkaProducer.WriteMessages(context.Background(), message)

	if err != nil {
		global.Logger.Error("Error while writing message to kafka", zap.Error(err))
		return response.ErrSendEmailOtp
	}

	return response.ErrCodeSuccess
}
