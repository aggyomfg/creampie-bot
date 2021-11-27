package store

// Store ...
type Store interface {
	Duel() DuelRepository
	HolidayToday() HolidayRepository
}
