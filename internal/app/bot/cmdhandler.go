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
	case "sticker_mode":
		if adminCheck(server, bot, msg) {
			botSkills.SwitchSkill(
				skills.Skill{
					Name:     "StickerMode",
					Function: skills.StickerMode(),
				},
			)
			sendMsg.Text = printCurrentSkills()
			break
		}
		sendMsg.Text = onlyAdminString

	case "translate_mode":
		if adminCheck(server, bot, msg) {
			botSkills.SwitchSkill(
				skills.Skill{
					Name: "TranslateMode",
					Function: skills.TranslateMode(
						server.config.YaTranslateAPIKey,
						server.config.Logger),
				},
			)
			sendMsg.Text = printCurrentSkills()
			break
		}
		sendMsg.Text = onlyAdminString

	default:
		sendMsg.Text = "Неизвестная команда :("
	}
	bot.Send(sendMsg)
}

func printCurrentSkills() string {
	message := "Текущие режимы:\n"
	names := botSkills.GetCurrentSkillsName()
	if len(names) > 0 {
		for _, name := range names {
			message = message + name + "\n"
		}
	} else {
		message = message + "Стандартный чат"
	}
	return message
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
