package response

const (
	ErrCodeSuccess      = 20001 // Success
	ErrCodeParamInvalid = 20002 // Email is invalid

	ErrInvalidToken = 30001
	ErrInvalOTP     = 30002
	ErrSendEmailOtp = 30003

	// Error Authenticate
	ErrCodeAuthFailed      = 40005
	ErrCodeAuthTokenFailed = 40001

	// Error Register
	ErrCodeUserAlreadyExists = 50002

	// Error Login
	ErrCodeOtpAlreadyExists = 60009
	ErrCodeOtpNotExists     = 60008

	ErrCodeTwoFactorAuthSetupFailed  = 80001
	ErrCodeTwoFactorAuthVerifyFailed = 80002
)

var msg = map[int]string{
	ErrCodeSuccess:      "Success",
	ErrCodeParamInvalid: "Email is invalid",

	ErrInvalidToken: "Token is invalid",
	ErrInvalOTP:     "OTP is invalid",
	ErrSendEmailOtp: "Send email OTP failed",

	ErrCodeAuthFailed:      "Authenticate Failed",
	ErrCodeAuthTokenFailed: "Unauthorized",

	ErrCodeUserAlreadyExists: "User already exists",

	ErrCodeOtpAlreadyExists: "OTP exists but not registered",
	ErrCodeOtpNotExists:     "OTP not exists",

	ErrCodeTwoFactorAuthSetupFailed:  "Two Factor Authentication Setup failed",
	ErrCodeTwoFactorAuthVerifyFailed: "Two Factor Authentication Verify failed",
}
