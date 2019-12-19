package skills

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// StickerMode ...
func StickerMode() func(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	return func(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
		if !(msg.Sticker != nil || msg.Animation != nil) {
			deleteMessageConfig := tgbotapi.DeleteMessageConfig{
				ChatID:    msg.Chat.ID,
				MessageID: msg.MessageID,
			}
			_, err := bot.DeleteMessage(deleteMessageConfig)

			if err != nil {
				log.Error(err)
			}
		}
	}
}
