package config

import "github.com/9ssi7/turnstile"

var TurnstileClient turnstile.Service

func InitTurnStileClient() {
	TurnstileClient = turnstile.New(turnstile.Config{
		Secret: TURNSTILE_SECRET_KEY,
	})
}
