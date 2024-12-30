package response

const (
	ErrCodeSuccess      = 20001 // Success
	ErrCodeParamInvalid = 20002 // Email is invalid

	ErrInvalidToken = 30001
	ErrInvalOTP     = 30002
	ErrSendEmailOtp = 30003

	ErrCodeUserAlreadyExists = 50002
)

var msg = map[int]string{
	ErrCodeSuccess:      "Success",
	ErrCodeParamInvalid: "Email is invalid",

	ErrInvalidToken: "Token is invalid",
	ErrInvalOTP:     "OTP is invalid",
	ErrSendEmailOtp: "Send email OTP failed",

	ErrCodeUserAlreadyExists: "User already exists",
}
