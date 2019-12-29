package memorystore

import (
	"github.com/aggyomfg/creampie-bot/internal/app/model"
	"github.com/aggyomfg/creampie-bot/internal/app/store"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// DuelRepository ...
type DuelRepository struct {
	store *Store
	duels map[int]*model.Duel
}

// Create ...
func (r *DuelRepository) Create(d *model.Duel) error {
	d.ID = len(r.duels) + 1
	r.duels[d.ID] = d

	return nil
}

// Find ...
func (r *DuelRepository) Find(id int) (*model.Duel, error) {
	u, ok := r.duels[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}

// FindByUser ...
func (r *DuelRepository) FindByUser(user tgbotapi.User) (*model.Duel, error) {
	for _, d := range r.duels {
		if d.Attacker == user {
			return d, nil
		}
	}

	return nil, store.ErrRecordNotFound
}
