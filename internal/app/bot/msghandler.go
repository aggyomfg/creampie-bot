package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func handleMsg(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {

	if msg.IsCommand() {
		handeCommand(bot, msg)
		return
	}
	for i, s := range botSkills.List {
		log.Debug(i, s)
		s.Function(bot, msg)
	}

}
