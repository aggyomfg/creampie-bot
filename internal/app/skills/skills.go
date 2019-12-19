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

// RegisterSkill ...
func (skills *Skills) RegisterSkill(skill Skill) []Skill {
	for _, s := range skills.List {
		if s.Name == skill.Name {
			return skills.List
		}
	}
	skills.List = append(skills.List, skill)
	return skills.List
}

// UnregisterSkill ...
func (skills *Skills) UnregisterSkill(skill Skill) []Skill {
	for p, s := range skills.List {
		if s.Name == skill.Name {
			skills.List = append(skills.List[:p], skills.List[p+1:]...)
			return skills.List
		}
	}
	return skills.List
}
