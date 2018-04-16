package dto

// A ResetPasswordInfo represents a request for resetting password
type ResetPasswordInfo struct {
	Password        string `json:"Password"`
	ConfirmPassword string `json:"ConfirmPassword"`
}
