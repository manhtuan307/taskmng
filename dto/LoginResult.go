package dto

// A LoginResult represents authentication result together with token
type LoginResult struct {
	IsSuccess   bool   `json:"isSuccess"`
	Message     string `json:"message"`
	UserID      string `json:"userId"`
	Token       string `json:"token"`
	ExpiredTime string `json:"expiredTime"`
}
