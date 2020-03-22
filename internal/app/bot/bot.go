package bot

import (
	"net/http"
	"net/url"
	"os"

	"github.com/aggyomfg/creampie-bot/internal/app/model"
	"github.com/aggyomfg/creampie-bot/internal/app/store/memorystore"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Start runs bot instance
func Start(config *model.Config) {
	logger := config.Logger
	tgbotapi.SetLogger(logger)

	store := memorystore.New()
	var client *http.Client
	if config.Proxy != "" {
		proxy, err := url.Parse(config.Proxy)
		if err != nil {
			logger.Error(err)
			os.Exit(1)
		}
		client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxy)}}
	} else {
		client = &http.Client{}
	}

	bot, err := tgbotapi.NewBotAPIWithClient(config.TelegramToken, client)
	// bot.Debug = true
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	srv := newServer(store, *bot, config)

	srv.Run()
}
