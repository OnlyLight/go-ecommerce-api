package impl

import (
	"context"
	"database/sql"
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
	"github.com/onlylight29/go-ecommerce-backend-api/internal/utils/auth"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/utils/crypto"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/utils/random"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/utils/sendto"
	"github.com/onlylight29/go-ecommerce-backend-api/pkg/response"
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

// Register -> verify_OTP -> change Pass -> login

// Implement interface of IUserLogin
func (s *sUserLogin) Login(ctx context.Context, in *model.LoginInput) (codeResult int, out model.LoginOutput, err error) {
	userBase, err := s.r.GetOneUserInfo(ctx, in.UserAccount)
	if err != nil {
		return response.ErrCodeAuthFailed, out, err
	}

	// 2. Check password
	if !crypto.MatchingPassword(userBase.UserPassword, in.UserPassword, userBase.UserSalt) {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("does not match password:: %v", err)
	}

	// 3. Check two-factor authentication
	isTwoFactorEnable, err := s.r.IsTwoFactorEnabled(ctx, uint32(userBase.UserID))
	if err != nil {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("does not match")
	}

	if isTwoFactorEnable > 0 {
		// 3.1 Send otp to in.TwoFactorEmail
		keyUserLoginTwoFactor := crypto.GetHash(utils.Get2FaKey(strconv.Itoa(int(userBase.UserID))))
		err = global.RDB.Set(ctx, keyUserLoginTwoFactor, "111111", time.Duration(constance.TIME_OTP_REGISTER)*time.Minute).Err()
		if err != nil {
			return response.ErrCodeAuthFailed, out, fmt.Errorf("set OTP register failed")
		}

		infoUserTwoFactor, err := s.r.GetTwoFactorMethodByIDAndType(ctx, database.GetTwoFactorMethodByIDAndTypeParams{
			UserID:            uint32(userBase.UserID),
			TwoFactorAuthType: database.PreGoAccUserTwoFactor9999TwoFactorAuthTypeEMAIL,
		})

		if err != nil {
			return response.ErrCodeAuthFailed, out, fmt.Errorf("get OTP register failed")
		}

		// send mail
		go sendto.SendTextEmailOtp([]string{infoUserTwoFactor.TwoFactorEmail.String}, constance.HOST_EMAIL, "111111")
		out.Message = "Send OTP 2FA to Email..."

		return response.ErrCodeSuccess, out, nil
	}

	// 4. Update password time with goroutine async
	go s.r.LoginUserBase(ctx, database.LoginUserBaseParams{
		UserLoginIp:  sql.NullString{String: "127.0.0.1", Valid: true},
		UserAccount:  in.UserAccount,
		UserPassword: in.UserPassword,
	})

	// 5. Create UUID User
	subToken := utils.GenerateCliTokenUUID(int(userBase.UserID))
	fmt.Println("subToken::", subToken)

	// 6. Get UserInfo table
	infoUser, err := s.r.GetUser(ctx, uint64(userBase.UserID))
	if err != nil {
		return response.ErrCodeAuthFailed, out, err
	}

	// 7. Convert to json in order to save to Redis
	infoUserJson, err := json.Marshal(infoUser)
	if err != nil {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("Convert to JSON failed:: %v", err)
	}

	// 8. Give infoUserJson to Redis with key = subToken
	err = global.RDB.Set(ctx, subToken, infoUserJson, time.Duration(constance.TIME_OTP_REGISTER)*time.Minute).Err()
	if err != nil {
		return response.ErrCodeAuthFailed, out, err
	}

	// 9. create token
	out.Token, err = auth.CreateToken(subToken)
	if err != nil {
		return
	}

	return response.ErrCodeSuccess, out, nil
}

func (s *sUserLogin) Register(ctx context.Context, in *model.RegisterInput) (codeResult int, err error) {
	// 1. Hash email
	fmt.Printf("VerifyKey:: %s\n", in.VerifyKey)
	fmt.Printf("VerifyType:: %d\n", in.VerifyType)
	fmt.Printf("Purpose:: %s\n", in.VerifyPurpose)
	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))
	fmt.Printf("hashKey:: %s\n", hashKey)

	// 2. Check user exist in user base
	userFound, err := s.r.CheckUserBaseExist(ctx, in.VerifyKey)
	if err != nil {
		global.Logger.Error("CheckUserBaseExist", zap.Error(err))
		return response.ErrCodeUserAlreadyExists, err
	}

	if userFound > 0 {
		return response.ErrCodeUserAlreadyExists, fmt.Errorf("User already exists")
	}

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
	if errCode, errMsg := utils.HandleGetKeyRedis(otpFound, err); errMsg != nil {
		return errCode, errMsg
	}

	if otpFound != "" {
		return response.ErrCodeOtpAlreadyExists, fmt.Errorf("OTP already exists")
	}

	//// 4. Generate OTP
	otpNew := random.GenerateSixDigitOTP()
	if in.VerifyPurpose == "TEST_USER" {
		otpNew = 123456
	}
	fmt.Println("OTP::", otpNew)

	//// 5. Save OTP to Redis
	err = global.RDB.SetEx(ctx, userKey, strconv.Itoa(otpNew), time.Duration(constance.TIME_OTP_REGISTER)*time.Minute).Err()
	if err != nil {
		return response.ErrInvalOTP, err
	}
	fmt.Println("Save OTP to Redis success!!")

	otpFound2, _ := global.RDB.Get(ctx, userKey).Result()
	fmt.Printf("userKey::%s\n", userKey)
	fmt.Printf("OTP found:: %v\n", otpFound2)

	//// 6. Send OTP to user
	switch in.VerifyType {
	case constance.EMAIL:
		// err = sendto.SendTextEmailOtp([]string{in.VerifyKey}, constance.HOST_EMAIL, strconv.Itoa(otpNew))
		// if err != nil {
		// 	return response.ErrSendEmailOtp, err
		// }

		// // 7. Send OTP to Kafka
		// body := make(map[string]interface{})
		// body["otp"] = otpNew
		// body["email"] = in

		// jsonBody, _ := json.Marshal(body)

		// message := kafka.Message{
		// 	Key:   []byte("otp-auth"),
		// 	Value: []byte(jsonBody),
		// 	Time:  time.Now(),
		// }

		// err = global.KafkaProducer.WriteMessages(context.Background(), message)

		// if err != nil {
		// 	global.Logger.Error("Error while writing message to kafka", zap.Error(err))
		// 	return response.ErrSendEmailOtp, err
		// }

		// // 7.2. Get OTP from Kafka

		// 8. Save to DB
		result, err := s.r.InsertOTPVerify(ctx, database.InsertOTPVerifyParams{
			VerifyOtp:     strconv.Itoa(otpNew),
			VerifyType:    1,
			VerifyKey:     in.VerifyKey,
			VerifyKeyHash: hashKey,
		})

		if err != nil {
			return response.ErrSendEmailOtp, err
		}

		// 9. Get LastID
		lastIdVerifyUser, err := result.LastInsertId()
		if err != nil {
			return response.ErrSendEmailOtp, err
		}

		fmt.Println("LastID::", lastIdVerifyUser)
		return response.ErrCodeSuccess, nil

	case constance.MOBILE:
		return response.ErrCodeSuccess, nil
	}

	return response.ErrCodeSuccess, nil
}

func (s *sUserLogin) VerifyOTP(ctx context.Context, in *model.VerifyInput) (out model.VerifyOTPOutput, err error) {
	// 1. Hash email
	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))
	userKey := utils.GetUserKey(hashKey)

	// 2. get OTP from Redis
	otpFound, err := global.RDB.Get(ctx, userKey).Result()
	if err != nil {
		return out, fmt.Errorf("err::otpFound::%v", err)
	}
	fmt.Printf("otpFound::%v", otpFound)

	err = utils.HandleOTPValidation(ctx, hashKey, otpFound, in.VerifyCode)
	if err != nil {
		return out, fmt.Errorf("err::HandleOTPValidation::%v", err)
	}

	infoOTP, err := s.r.GetInfoOTP(ctx, hashKey)
	if err != nil {
		return out, fmt.Errorf("err::infoOTP::%v", err)
	}
	fmt.Printf("infoOTP::%v", infoOTP)

	// Update status verified
	err = s.r.UpdateUserVerificationStatus(ctx, hashKey)
	if err != nil {
		return out, fmt.Errorf("err::UpdateUserVerificationStatus::%v", err)
	}

	// Output
	out.Token = infoOTP.VerifyKeyHash
	out.Message = "Success"

	return out, nil
}

func (s *sUserLogin) UpdatePasswordRegister(ctx context.Context, token string, password string) (userId int, err error) {
	//1. Check Token is already verified --- user_verify table
	infoOTP, err := s.r.GetInfoOTP(ctx, token)
	if err != nil {
		return response.ErrCodeOtpNotExists, err
	}

	// 2. check OTP is verified
	if !infoOTP.IsVerified {
		return response.ErrCodeOtpNotExists, fmt.Errorf("OTP not verified")
	}

	// 2.1 check token is exist on user_base table

	// 3. Update user_base password
	userBase := database.AddUserBaseParams{}
	userBase.UserAccount = infoOTP.VerifyKey

	userSalt, err := crypto.GenerateSalt(16)
	if err != nil {
		return response.ErrCodeOtpNotExists, err
	}

	userBase.UserSalt = userSalt
	userBase.UserPassword = crypto.GetHash(password + userSalt)

	// 4. add userBase to user_base table
	newUserBase, err := s.r.AddUserBase(ctx, userBase)
	if err != nil {
		return response.ErrCodeOtpNotExists, err
	}

	user_id, err := newUserBase.LastInsertId()
	if err != nil {
		return response.ErrCodeOtpNotExists, err
	}

	// add user_id to user_info table
	newUserInfo, err := s.r.AddUserAutoUserId(ctx, database.AddUserAutoUserIdParams{
		UserID:               uint64(user_id),
		UserAccount:          infoOTP.VerifyKey,
		UserNickname:         sql.NullString{String: infoOTP.VerifyKey, Valid: true},
		UserAvatar:           sql.NullString{String: "", Valid: true},
		UserState:            1,
		UserMobile:           sql.NullString{String: "", Valid: true},
		UserGender:           sql.NullInt16{Int16: 0, Valid: true},
		UserBirthday:         sql.NullTime{Time: time.Time{}, Valid: false},
		UserEmail:            sql.NullString{String: infoOTP.VerifyKey, Valid: true},
		UserIsAuthentication: 1,
	})

	if err != nil {
		return response.ErrCodeOtpNotExists, err
	}

	user_Id, err := newUserInfo.LastInsertId()
	if err != nil {
		return response.ErrCodeOtpNotExists, err
	}

	return int(user_Id), nil
}

func (s *sUserLogin) IsTwoFactorEnabled(ctx context.Context, userId int) (codeResult int, rs bool, err error) {
	return response.ErrCodeSuccess, true, nil
}

func (s *sUserLogin) SetupTwoFactorAuth(ctx context.Context, in *model.SetupTwoFactorAuthInput) (codeResult int, err error) {
	// 1. check isTwoFactorEnabled -> true return
	isTrueFactorAuth, err := s.r.IsTwoFactorEnabled(ctx, in.UserId)
	if err != nil {
		return response.ErrCodeTwoFactorAuthSetupFailed, err
	}

	if isTrueFactorAuth > 0 {
		return response.ErrCodeTwoFactorAuthSetupFailed, fmt.Errorf("two-factor authentication is already enabled")
	}

	// 2. Create new record type Auth
	err = s.r.EnableTwoFactorTypeEmail(ctx, database.EnableTwoFactorTypeEmailParams{
		UserID:            in.UserId,
		TwoFactorAuthType: database.PreGoAccUserTwoFactor9999TwoFactorAuthTypeEMAIL,
		TwoFactorEmail:    sql.NullString{String: in.TwoFactorEmail, Valid: true},
	})

	if err != nil {
		return response.ErrCodeTwoFactorAuthSetupFailed, err
	}

	// 3. send otp to in.TwoFactorEmail
	hashUserTwoFactor := crypto.GetHash(strconv.Itoa(int(in.UserId)))
	keyUserTwoFactor := utils.Get2FaKey(hashUserTwoFactor)

	err = global.RDB.Set(ctx, keyUserTwoFactor, "111111", time.Duration(constance.TIME_OTP_REGISTER)*time.Minute).Err()
	if err != nil {
		return response.ErrCodeTwoFactorAuthSetupFailed, err
	}

	return response.ErrCodeSuccess, nil
}

func (s *sUserLogin) VerifyTwoFactorAuth(ctx context.Context, in *model.TwoFactorVerificationInput) (codeResult int, err error) {
	isTwoFactorAuth, err := s.r.IsTwoFactorEnabled(ctx, in.UserId)
	if err != nil {
		return response.ErrCodeTwoFactorAuthVerifyFailed, err
	}

	if isTwoFactorAuth > 0 {
		return response.ErrCodeTwoFactorAuthVerifyFailed, fmt.Errorf("2FA authentication is not enabled")
	}

	// 2. Check OTP in redis avalable
	hashUserTwoFactor := crypto.GetHash(strconv.Itoa(int(in.UserId)))
	keyUserTwoFactor := utils.Get2FaKey(hashUserTwoFactor)

	otpVerifyAuth, err := global.RDB.Get(ctx, keyUserTwoFactor).Result()

	if errCode, errMsg := utils.HandleGetKeyRedis(otpVerifyAuth, err); errMsg != nil {
		return errCode, errMsg
	}

	// 3. Check OTP
	if otpVerifyAuth != in.TwoFactorCode {
		return response.ErrCodeTwoFactorAuthVerifyFailed, fmt.Errorf("OTP does not match")
	}

	// 4. Update status
	err = s.r.UpdateTwoFactorStatus(ctx, database.UpdateTwoFactorStatusParams{
		UserID:            in.UserId,
		TwoFactorAuthType: database.PreGoAccUserTwoFactor9999TwoFactorAuthTypeEMAIL,
	})

	if err != nil {
		return response.ErrCodeTwoFactorAuthVerifyFailed, err
	}

	// 5. Remove OTP
	_, err = global.RDB.Del(ctx, keyUserTwoFactor).Result()
	if err != nil {
		return response.ErrCodeTwoFactorAuthVerifyFailed, err
	}

	return response.ErrCodeSuccess, nil
}
