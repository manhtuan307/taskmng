package dto

// A RegistrationInfo represents registration info
type RegistrationInfo struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}
