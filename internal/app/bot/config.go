package bot

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Config ...
type Config struct {
	BindAddr      string `env:"BIND_ADDR"`
	TelegramToken string `env:"TELEGRAM_TOKEN"`
	Proxy         string `env:"PROXY_URL"`
	Logger        *logrus.Logger
}

// NewConfig ...
func NewConfig() *Config {
	level, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logrus.Fatal("Cant parse LOG_LEVEL")
	}

	var log = &logrus.Logger{
		Out:       os.Stderr,
		Formatter: new(logrus.TextFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     level,
	}

	return &Config{
		Logger: log,
	}
}
