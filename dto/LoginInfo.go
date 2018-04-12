package dto

// A LoginInfo represents authentication information which user provides
type LoginInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
