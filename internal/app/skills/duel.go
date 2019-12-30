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
	var currentDuel *model.Duel
	if _, err := store.Duel().FindByUser(*attacker); err != nil {
		if _, err := store.Duel().FindByUser(*defender); err != nil {
			currentDuel, _ = createDuel(attacker, defender, store)
		} else {
			sendMsg.Text = defender.String() + " уже учавствует в дуэли :("
			bot.Send(sendMsg)
			return
		}
	} else {
		sendMsg.Text = attacker.String() + " уже учавствует в дуэли :("
		bot.Send(sendMsg)
		return
	}
	sendMsg.Text = attacker.String() + " vs. " + defender.String() + " Duel id is: " + strconv.Itoa(currentDuel.ID)
	bot.Send(sendMsg)
}

func createDuel(attacker *tgbotapi.User, defender *tgbotapi.User, store store.Store) (*model.Duel, error) {
	store.Duel().Create(&model.Duel{
		Attacker: *attacker,
		Defender: *defender,
	})
	return store.Duel().GetLast()
}
