package dto

// A ChangePasswordRequest represents a request for changing password
type ChangePasswordRequest struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}
