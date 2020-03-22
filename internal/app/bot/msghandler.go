package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func handleMsg(server *Server, bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {

	if msg.IsCommand() {
		handeCommand(server, bot, msg)
		return
	}
	for i, s := range botSkills.List {
		server.log.Debug(i, s)
		s.Function(bot, msg)
	}

}
