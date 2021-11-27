package model

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// HolidayToday ...
type HolidayToday struct {
	ID      int
	User    tgbotapi.User
	Holiday string
}
