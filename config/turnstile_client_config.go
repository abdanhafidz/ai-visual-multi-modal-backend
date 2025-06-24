package config

import (
	"github.com/meyskens/go-turnstile"
)

var TurnstileClient *turnstile.Turnstile

func InitTurnStileClient() {
	TurnstileClient = turnstile.New(TURNSTILE_SECRET_KEY)
}
