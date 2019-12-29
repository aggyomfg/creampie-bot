package skills

import (
	"strconv"

	"github.com/aggyomfg/creampie-bot/internal/app/model"
	"github.com/aggyomfg/creampie-bot/internal/app/store"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Duel ...
func Duel(bot *tgbotapi.BotAPI, msg *tgbotapi.Message, store store.Store) {
	sendMsg := tgbotapi.NewMessage(msg.Chat.ID, "")
	if msg.ReplyToMessage == nil {
		sendMsg.Text = "Нужно ответить на сообщение!"
		bot.Send(sendMsg)
		return
	}
	defender := msg.ReplyToMessage.From
	attacker := msg.From
	store.Duel().Create(&model.Duel{
		ID:       1,
		Attacker: *msg.From,
		Defender: *msg.ReplyToMessage.From,
	})
	currentDuel, _ := store.Duel().Find(1)
	sendMsg.Text = attacker.String() + " vs. " + defender.String() + " Duel id is: " + strconv.Itoa(currentDuel.ID)
	bot.Send(sendMsg)
}
