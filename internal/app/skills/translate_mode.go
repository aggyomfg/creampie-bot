package skills

import (
	"math/rand"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"

	translate "github.com/dafanasev/go-yandex-translate"
	"golang.org/x/text/language"
)

func chooseRandomLanguage() string {
	rand.Seed(time.Now().Unix())
	languages := []string{
		language.Ukrainian.String(),
		language.Bulgarian.String(),
		"be", // Белорусский
	}
	n := rand.Int() % len(languages)
	return languages[n]
}

func getCountyFlag(lang string) string {
	switch lang {
	case language.Ukrainian.String():
		return "🇺🇦"
	case language.Bulgarian.String():
		return "🇧🇬"
	case "be":
		return "🇧🇾"
	default:
		return "🏳️‍🌈"
	}
}

// TranslateMode ...
func TranslateMode(apiKey string, log *logrus.Logger) func(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	return func(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
		deleteMessage(bot, msg)
		defaultMessage := tgbotapi.NewMessage(msg.Chat.ID, getCountyFlag("none")+" "+msg.From.String()+": ")
		lang := chooseRandomLanguage()
		if !(msg.Sticker != nil || msg.Animation != nil) {
			sendMsg := tgbotapi.NewMessage(msg.Chat.ID, "")
			tr := translate.New(apiKey)
			translation, err := tr.Translate(lang, msg.Text)
			if err != nil {
				log.Error(err)
			} else {
				sendMsg.Text = getCountyFlag(lang) + " " + msg.From.String() + ": " + translation.Result()
			}
			bot.Send(sendMsg)
		}
		if msg.Sticker != nil {
			sendTextMsg := defaultMessage
			sendStickerMsg := tgbotapi.NewStickerShare(msg.Chat.ID, msg.Sticker.FileID)
			bot.Send(sendTextMsg)
			bot.Send(sendStickerMsg)
		}
		if msg.Animation != nil {
			sendTextMsg := defaultMessage
			sendAnimationMsg := tgbotapi.NewAnimationShare(msg.Chat.ID, msg.Animation.FileID)
			bot.Send(sendTextMsg)
			bot.Send(sendAnimationMsg)
		}

	}
}
