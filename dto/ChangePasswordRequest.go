package dto

// A ChangePasswordRequest represents a request for changing password
type ChangePasswordRequest struct {
	OldPassword     string `json:"OldPassword"`
	Password        string `json:"Password"`
	ConfirmPassword string `json:"ConfirmPassword"`
}
