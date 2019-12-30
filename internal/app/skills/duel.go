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
	if defender.ID == bot.Self.ID {
		sendMsg.Text = "Ты на кого батон крошишь? 😡"
		bot.Send(sendMsg)
		return
	}
	if u1, _, d := checkDuelers(attacker, defender, store); d == nil {
		duel, _ := createDuel(attacker, defender, store)
		sendMsg.Text = "@" + attacker.String() + " против " + "@" + defender.String() + " \n🔫 на столе, приготовтесь\n" + "Первый ход за: @" + duel.CurrentPlayer.String() + "\n/shot"
	} else {
		sendMsg.Text = "@" + u1.String() + " уже учавствует в дуэли № " + strconv.Itoa(d.ID) + " против " + "@" + getDuelEnemy(u1, d).String()
	}
	bot.Send(sendMsg)
}

// DuelShot ...
func DuelShot(bot *tgbotapi.BotAPI, msg *tgbotapi.Message, store store.Store) {
	sendMsg := tgbotapi.NewMessage(msg.Chat.ID, "")
	if d, _ := store.Duel().FindByUser(*msg.From); d != nil {
		currentBarrel := diceIt()
		if d.CurrentPlayer.ID == msg.From.ID {
			sendMsg.Text = "@" + msg.From.String() + " крутит барабан, жмёт на спуск и остаётся жив!\n😤😤😤\n" + "Следующий ходит: @" + getDuelEnemy(msg.From, d).String() + "\n/shot"
			d.CurrentPlayer = *getDuelEnemy(msg.From, d)
			if d.BarrelRound == currentBarrel {
				sendMsg.Text = "@" + msg.From.String() + " крутит барабан, жмёт на спуск и раскидывает мозги по комнате! ☠️☠️☠️\n@" + getDuelEnemy(msg.From, d).String() + " победил!"
				store.Duel().Delete(d.ID)
			}
		} else {
			sendMsg.Text = "🤬 Положи револьвер! Сейчас ходит @" + d.CurrentPlayer.String()
		}
		bot.Send(sendMsg)
		deleteMessage(bot, msg)
		return
	}
	sendMsg.ReplyToMessageID = msg.MessageID
	sendMsg.Text = "Ты не учавствуешь в дуэли!"
	bot.Send(sendMsg)
}

func createDuel(attacker *tgbotapi.User, defender *tgbotapi.User, store store.Store) (*model.Duel, error) {
	var currentPlayer *tgbotapi.User
	if diceIt() > 3 {
		currentPlayer = defender
	} else {
		currentPlayer = attacker
	}
	store.Duel().Create(&model.Duel{
		Attacker:      *attacker,
		Defender:      *defender,
		BarrelRound:   diceIt(),
		CurrentPlayer: *currentPlayer,
	})
	return store.Duel().GetLast()
}

func getDuelEnemy(user *tgbotapi.User, duel *model.Duel) *tgbotapi.User {
	if user.ID == duel.Attacker.ID {
		return &duel.Defender
	}
	return &duel.Attacker
}

func checkDuelers(attacker *tgbotapi.User, defender *tgbotapi.User, store store.Store) (*tgbotapi.User, *tgbotapi.User, *model.Duel) {
	var (
		d   *model.Duel
		err error
	)
	if d, err = store.Duel().FindByUser(*attacker); err != nil {
		if d, err = store.Duel().FindByUser(*defender); err != nil {
			return nil, nil, nil
		}
		return defender, attacker, d
	}
	return attacker, defender, d
}
