package models

type LoginRequest struct {
	PassPhrase string `json:"pass_phrase" binding:"required"`
	TurnStile  string `json:"turnstile_payload" binding:"required"`
	IPAddress  string `json:"ip_address"`
}
