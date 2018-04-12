package dto

// A LoginResult represents authentication result together with token
type LoginResult struct {
	IsSuccess   bool   `json:"IsSuccess"`
	Message     string `json:"Message"`
	UserID      string `json:"UserId"`
	Token       string `json:"Token"`
	ExpiredTime string `json:"ExpiredTime"`
}
