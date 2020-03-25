package skills

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

var (
	log *logrus.Logger
)

// Skills ...
type Skills struct {
	List []Skill
}

// Skill ...
type Skill struct {
	Name     string
	Function func(bot *tgbotapi.BotAPI, msg *tgbotapi.Message)
}

// GetCurrentSkillsName ...
func (skills *Skills) GetCurrentSkillsName() []string {
	var names []string
	for _, s := range skills.List {
		names = append(names, s.Name)
	}
	return names
}

// SwitchSkill switches skill on/off
func (skills *Skills) SwitchSkill(skill Skill) []Skill {
	for p, s := range skills.List {
		if s.Name == skill.Name {
			skills.List = append(skills.List[:p], skills.List[p+1:]...)
			return skills.List
		}
	}
	skills.List = append(skills.List, skill)
	return skills.List
}

func deleteMessage(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	deleteMessageConfig := tgbotapi.DeleteMessageConfig{
		ChatID:    msg.Chat.ID,
		MessageID: msg.MessageID,
	}
	_, err := bot.DeleteMessage(deleteMessageConfig)
	if err != nil {
		log.Error(err)
	}
}
