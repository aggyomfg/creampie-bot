package skills

import (
	"math/rand"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"

	translate "github.com/dafanasev/go-yandex-translate"
	"golang.org/x/text/language"
)

func chooseRandomLanguage() language.Tag {
	rand.Seed(time.Now().Unix())
	languages := []language.Tag{
		language.Ukrainian,
		language.Polish,
	}
	n := rand.Int() % len(languages)
	return languages[n]
}

func getCountyFlag(lang language.Tag) string {
	switch lang {
	case language.Ukrainian:
		return "🇺🇦"
	case language.Polish:
		return "🇵🇱"
	default:
		return "🏳️‍🌈"
	}
}

// TranslateMode ...
func TranslateMode(apiKey string, log *logrus.Logger) func(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	return func(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
		deleteMessage(bot, msg)
		sendMsg := tgbotapi.NewMessage(msg.Chat.ID, "")
		tr := translate.New(apiKey)
		lang := chooseRandomLanguage()
		translation, err := tr.Translate(lang.String(), msg.Text)
		if err != nil {
			log.Error(err)
		} else {
			sendMsg.Text = getCountyFlag(lang) + " " + msg.From.String() + ": " + translation.Result()
		}
		bot.Send(sendMsg)
	}
}
