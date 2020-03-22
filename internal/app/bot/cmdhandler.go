package bot

import (
	"github.com/aggyomfg/creampie-bot/internal/app/skills"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func handeCommand(server *Server, bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	onlyAdminString := "Только админ может менять режим"
	sendMsg := tgbotapi.NewMessage(msg.Chat.ID, "")
	switch msg.Command() {
	case "help":
		sendMsg.Text = "type /sayhi or /status."
	case "sayhi":
		sendMsg.Text = "Hi :)"
	case "status":
		sendMsg.Text = "I'm ok."
	case "duel":
		skills.Duel(bot, msg, server.store)
	case "shot":
		skills.DuelShot(bot, msg, server.store)
	case "dice":
		skills.RollDice(bot, msg)
	case "sticker_mode_on":
		if adminCheck(server, bot, msg) {
			botSkills.RegisterSkill(
				skills.Skill{
					Name:     "StickerMode",
					Function: skills.StickerMode(),
				},
			)
			sendMsg.Text = "Общаемся только стикерами и гифками :P"
			break
		}
		sendMsg.Text = onlyAdminString

	case "sticker_mode_off":
		if adminCheck(server, bot, msg) {
			botSkills.UnregisterSkill(
				skills.Skill{
					Name: "StickerMode",
				},
			)
			sendMsg.Text = "Общаемся нормально!"
			break
		}
		sendMsg.Text = onlyAdminString
	case "translate_mode_on":
		if adminCheck(server, bot, msg) {
			botSkills.RegisterSkill(
				skills.Skill{
					Name: "TranslateMode",
					Function: skills.TranslateMode(
						server.config.YaTranslateAPIKey,
						server.config.Logger),
				},
			)
			sendMsg.Text = "Побалакаем?"
			break
		}
		sendMsg.Text = onlyAdminString

	case "translate_mode_off":
		if adminCheck(server, bot, msg) {
			botSkills.UnregisterSkill(
				skills.Skill{
					Name: "TranslateMode",
				},
			)
			sendMsg.Text = "Общаемся нормально!"
			break
		}
		sendMsg.Text = onlyAdminString
	default:
		sendMsg.Text = "Неизвестная команда :("
	}
	bot.Send(sendMsg)
}

func adminCheck(server *Server, bot *tgbotapi.BotAPI, msg *tgbotapi.Message) bool {
	admins, err := bot.GetChatAdministrators(tgbotapi.ChatConfig{
		ChatID: msg.Chat.ID,
	})
	if err != nil {
		server.log.Error(err)
	}
	for _, a := range admins {
		if msg.From.ID == a.User.ID {
			return true
		}
	}
	return false
}
