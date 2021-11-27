package store

import (
	"github.com/aggyomfg/creampie-bot/internal/app/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// DuelRepository ...
type DuelRepository interface {
	Create(*model.Duel) error
	Find(int) (*model.Duel, error)
	Delete(int) error
	FindByUser(tgbotapi.User) (*model.Duel, error)
	GetLast() (*model.Duel, error)
}

type HolidayRepository interface {
	Create(*model.HolidayToday) error
	Find(int) (*model.HolidayToday, error)
	Delete(int) error
	FindByUser(tgbotapi.User) (string, error)
	GetLast() (*model.HolidayToday, error)
	GetAllHolidaysToday() ([]string, error)
	SetAllHolidaysToday([]string)
}
