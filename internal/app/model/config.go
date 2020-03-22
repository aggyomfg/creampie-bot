package model

import "github.com/sirupsen/logrus"

// Config ...
type Config struct {
	BindAddr          string `env:"BIND_ADDR"`
	TelegramToken     string `env:"TELEGRAM_TOKEN"`
	Proxy             string `env:"PROXY_URL"`
	YaTranslateAPIKey string `env:"YA_TRANSLATE_API_KEY"`
	Logger            *logrus.Logger
}
