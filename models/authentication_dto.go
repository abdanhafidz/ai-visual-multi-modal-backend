package models

type LoginRequest struct {
	PassPhrase string `json:"pass_phrase binding:"required"`
	TurnStile  string `json:"turnstile_payload binding:"required"`
}
