package service

import (
	"dev11/internal/domain"
	"time"
)

// Интерфейс сервиса календаря
type EventService interface {
	// Создание события
	Create(event domain.Event) (int, error)
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