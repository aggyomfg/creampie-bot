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
	case "sticker_mode":
		botSkills.RegisterSkill(
			skills.Skill{
				Name:     "StickerMode",
				Function: skills.StickerMode(),
			},
		)
		sendMsg.Text = "Общаемся только стикерами и гифками :P"
	case "sticker_mode_off":
		botSkills.UnregisterSkill(
			skills.Skill{
				Name: "StickerMode",
			},
		)
		sendMsg.Text = "Общаемся нормально!"
	default:
		sendMsg.Text = "I don't know that command"
	}
	bot.Send(sendMsg)
}
