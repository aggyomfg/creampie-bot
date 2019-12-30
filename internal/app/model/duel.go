package model

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Duel ...
type Duel struct {
	ID            int
	Attacker      tgbotapi.User
	Defender      tgbotapi.User
	BarrelRound   int
	CurrentPlayer tgbotapi.User
}
