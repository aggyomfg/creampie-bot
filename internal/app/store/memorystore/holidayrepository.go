package memorystore

import (
	"time"

	"github.com/aggyomfg/creampie-bot/internal/app/model"
	"github.com/aggyomfg/creampie-bot/internal/app/store"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// HolidayRepository ...
type HolidayRepository struct {
	store            *Store
	holidays         map[int]*model.HolidayToday
	todayHolidayList []string
	lastCheckTime    time.Time
}

// Create ...
func (r *HolidayRepository) Create(h *model.HolidayToday) error {
	h.ID = len(r.holidays) + 1
	r.holidays[h.ID] = h

	return nil
}

// GetLast ...
func (r *HolidayRepository) GetLast() (*model.HolidayToday, error) {
	for _, d := range r.holidays {
		if d.ID == len(r.holidays) {
			return d, nil
		}
	}
	return nil, store.ErrRecordNotFound
}

// Find ...
func (r *HolidayRepository) Find(id int) (*model.HolidayToday, error) {
	d, ok := r.holidays[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return d, nil
}

// Delete ...
func (r *HolidayRepository) Delete(id int) error {
	d, ok := r.holidays[id]
	if !ok {
		return store.ErrRecordNotFound
	}
	delete(r.holidays, d.ID)
	return nil
}

// FindByUser ...
func (r *HolidayRepository) FindByUser(user tgbotapi.User) (string, error) {
	for _, h := range r.holidays {
		if h.User == user {
			return h.Holiday, nil
		}
	}
	return "", store.ErrRecordNotFound
}

// GetAllHolidaysToday
func (r *HolidayRepository) GetAllHolidaysToday() ([]string, error) {
	if len(r.todayHolidayList) > 0 {
		return r.todayHolidayList, nil
	}
	return nil, store.ErrRecordNotFound
}

// SetAllHolidaysToday
func (r *HolidayRepository) SetAllHolidaysToday(holidaysToday []string) {
	r.todayHolidayList = holidaysToday
}

// UpdateLastCheckTime
func (r *HolidayRepository) UpdateLastCheckTime() {
	r.lastCheckTime = time.Now()
}

// GetLastCheckTime
func (r *HolidayRepository) GetLastCheckTime() time.Time {
	return r.lastCheckTime
}
