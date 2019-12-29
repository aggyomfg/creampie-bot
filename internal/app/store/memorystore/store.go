package memorystore

import (
	"github.com/aggyomfg/creampie-bot/internal/app/model"
	"github.com/aggyomfg/creampie-bot/internal/app/store"
)

// Store ...
type Store struct {
	duelRepository *DuelRepository
}

// New ...
func New() *Store {
	return &Store{}
}

// Duel ...
func (s *Store) Duel() store.DuelRepository {
	if s.duelRepository != nil {
		return s.duelRepository
	}

	s.duelRepository = &DuelRepository{
		store: s,
		duels: make(map[int]*model.Duel),
	}

	return s.duelRepository
}
