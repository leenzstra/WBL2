package repository

import (
	"dev11/internal/domain"
	"time"
)

// Интерфейс репозитория календаря
type EventRepository interface {
	// Создание события
	Create(domain.Event) (int, error)
	// Обновление события
	Update(int, domain.Event) error
	// Удаление события
	Delete(int) error
	// Получение событий для дня
	EventsForDay(time.Time) ([]domain.Event, error)
	// Получение событий для недели
	EventsForWeek(time.Time) ([]domain.Event, error)
	// Получение событий для месяца
	EventsForMonth(time.Time) ([]domain.Event, error)
}
