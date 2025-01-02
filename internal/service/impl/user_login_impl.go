package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/onlylight29/go-ecommerce-backend-api/global"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/constance"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/database"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/model"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/utils"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/utils/crypto"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/utils/random"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/utils/sendto"
	"github.com/onlylight29/go-ecommerce-backend-api/pkg/response"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type sUserLogin struct {
	r *database.Queries
}

func NewUserLoginImpl(r *database.Queries) *sUserLogin {
	return &sUserLogin{
		r: r,
	}
}

// Implement interface of IUserLogin
func (s *sUserLogin) Login(ctx context.Context) error {
	panic("unimplement yet")
}

func (s *sUserLogin) Register(ctx context.Context, in *model.RegisterInput) (codeResult int, error error) {
	// 1. Hash email
	fmt.Printf("VerifyKey:: %s\n", in.VerifyKey)
	fmt.Printf("VerifyType:: %d\n", in.VerifyType)
	fmt.Printf("Purpose:: %s\n", in.VerifyPurpose)
	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))
	fmt.Printf("hashKey:: %s\n", hashKey)

	//// 2. Check user exist in user base
	// userFound, err := s.r.CheckUserBaseExist(ctx, in.VerifyKey)
	// if err != nil {
	// 	global.Logger.Error("CheckUserBaseExist", zap.Error(err))
	// 	return response.ErrCodeUserAlreadyExists, err
	// }

	// if userFound > 0 {
	// 	return response.ErrCodeUserAlreadyExists, fmt.Errorf("User already exists")
	// }

	//// 3. Verify OTP
	userKey := utils.GetUserKey(hashKey)
	otpFound, err := global.RDB.Get(ctx, userKey).Result()

	// if err != nil {
	// 	if err != redis.Nil {
	// 		global.Logger.Error("Get OTP failed", zap.Error(err))
	// 		return response.ErrInvalOTP, err
	// 	} else {
	// 		global.Logger.Info("OTP not found")
	// 	}
	// }
	// if otpFound != "" {
	// 	return response.ErrCodeOtpAlreadyExists, fmt.Errorf("OTP already exists")
	// }
	if errCode, errMsg := utils.GetOtpFromRedis(otpFound, err); errMsg != nil {
		return errCode, errMsg
	}

	//// 4. Generate OTP
	otpNew := random.GenerateSixDigitOTP()
	if in.VerifyPurpose == "TEST_USER" {
		otpNew = 123456
	}
	fmt.Println("OTP::", otpNew)

	//// 5. Save OTP to Redis
	err = global.RDB.SetEx(ctx, hashKey, otpNew, time.Duration(constance.TIME_OTP_REGISTER)*time.Minute).Err()
	if err != nil {
		return response.ErrInvalOTP, err
	}

	//// 6. Send OTP to user
	switch in.VerifyType {
	case constance.EMAIL:
		err = sendto.SendTextEmailOtp([]string{in.VerifyKey}, constance.HOST_EMAIL, strconv.Itoa(otpNew))
		if err != nil {
			return response.ErrSendEmailOtp, err
		}

		//// 7. Send OTP to Kafka
		body := make(map[string]interface{})
		body["otp"] = otpNew
		body["email"] = in

		jsonBody, _ := json.Marshal(body)

		message := kafka.Message{
			Key:   []byte("otp-auth"),
			Value: []byte(jsonBody),
			Time:  time.Now(),
		}

		err = global.KafkaProducer.WriteMessages(context.Background(), message)

		if err != nil {
			global.Logger.Error("Error while writing message to kafka", zap.Error(err))
			return response.ErrSendEmailOtp, err
		}

		// // 7.2. Get OTP from Kafka

		// // 8. Save to DB
		// result, err := s.r.InsertOTPVerify(ctx, database.InsertOTPVerifyParams{
		// 	VerifyOtp:     strconv.Itoa(otpNew),
		// 	VerifyType:    sql.NullInt32{Int32: 1, Valid: true},
		// 	VerifyKey:     in.VerifyKey,
		// 	VerifyKeyHash: hashKey,
		// })

		// if err != nil {
		// 	return response.ErrSendEmailOtp, err
		// }

		// // 9. Get LastID
		// lastIdVerifyUser, err := result.LastInsertId()
		// if err != nil {
		// 	return response.ErrSendEmailOtp, err
		// }

		// fmt.Println("LastID::", lastIdVerifyUser)
		return response.ErrCodeSuccess, nil

	case constance.MOBILE:
		return response.ErrCodeSuccess, nil
	}

	return response.ErrCodeSuccess, nil
}

func (s *sUserLogin) VerifyOTP(ctx context.Context) error {
	panic("unimplement yet")
}

func (s *sUserLogin) UpdatePasswordRegister(ctx context.Context) error {
	panic("unimplement yet")
}
