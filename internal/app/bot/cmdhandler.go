package bot

import (
	"github.com/aggyomfg/creampie-bot/internal/app/skills"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func handeCommand(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	sendMsg := tgbotapi.NewMessage(msg.Chat.ID, "")
	switch msg.Command() {
	case "help":
		sendMsg.Text = "type /sayhi or /status."
	case "sayhi":
		sendMsg.Text = "Hi :)"
	case "status":
		sendMsg.Text = "I'm ok."
	case "dice":
		skills.RollDice(bot, msg)
	case "sticker_mode_on":
		if adminCheck(bot, msg) {
			botSkills.RegisterSkill(
				skills.Skill{
					Name:     "StickerMode",
					Function: skills.StickerMode(),
				},
			)
			sendMsg.Text = "Общаемся только стикерами и гифками :P"
			break
		}
		sendMsg.Text = "Только админ может менять режим"

	case "sticker_mode_off":
		if adminCheck(bot, msg) {
			botSkills.UnregisterSkill(
				skills.Skill{
					Name: "StickerMode",
				},
			)
			sendMsg.Text = "Общаемся нормально!"
			break
		}
		sendMsg.Text = "Только админ может менять режим"
	default:
		sendMsg.Text = "I don't know that command"
	}
	bot.Send(sendMsg)
}

func adminCheck(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) bool {
	admins, err := bot.GetChatAdministrators(tgbotapi.ChatConfig{
		ChatID: msg.Chat.ID,
	})
	if err != nil {
		log.Error(err)
	}
	for _, a := range admins {
		if msg.From.ID == a.User.ID {
			return true
		}
	}
	return false
}
