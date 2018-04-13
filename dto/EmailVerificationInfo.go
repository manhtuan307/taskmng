package dto

// A EmailVerificationInfo represents email verification info
type EmailVerificationInfo struct {
	Email      string `json:"Email"`
	VerifyCode string `json:"VerifyCode"`
}
