package bot

import (
	"github.com/aggyomfg/creampie-bot/internal/app/skills"
	"github.com/aggyomfg/creampie-bot/internal/app/store"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

var (
	botSkills = skills.Skills{}
)

type server struct {
	log   *logrus.Logger
	store store.Store
	bot   tgbotapi.BotAPI
}

func newServer(store store.Store, bot tgbotapi.BotAPI, logger *logrus.Logger) *server {
	s := &server{
		store: store,
		bot:   bot,
		log:   logger,
	}
	return s
}

func (s *server) Run() error {
	s.log.Printf("Authorized on account %s", s.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := s.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		s.log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		handleMsg(s, &s.bot, update.Message)
	}
	return nil
}
