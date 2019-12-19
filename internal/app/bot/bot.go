package bot

import (
	"github.com/aggyomfg/creampie-bot/internal/app/skills"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

var (
	log       *logrus.Logger
	botSkills = skills.Skills{}
)

// Start runs bot instance
func Start(config *Config) {

	log = config.Logger
	tgbotapi.SetLogger(log)

	bot, err := tgbotapi.NewBotAPI(config.TelegramToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		handleMsg(bot, update.Message)
	}
}
