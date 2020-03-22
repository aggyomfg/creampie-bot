package bot

import (
	"os"

	"github.com/aggyomfg/creampie-bot/internal/app/model"
	"github.com/sirupsen/logrus"
)

// NewConfig ...
func NewConfig() *model.Config {
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

	return &model.Config{
		Logger: log,
	}
}
