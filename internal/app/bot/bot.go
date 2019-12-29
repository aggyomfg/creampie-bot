package bot

import (
	"github.com/aggyomfg/creampie-bot/internal/app/store/memorystore"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"os"
)

// Start runs bot instance
func Start(config *Config) {
	logger := config.Logger
	tgbotapi.SetLogger(logger)

	store := memorystore.New()
	bot, err := tgbotapi.NewBotAPI(config.TelegramToken)
	// bot.Debug = true
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	srv := newServer(store, *bot, logger)
	srv.Run()
}
