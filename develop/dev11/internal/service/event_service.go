package service

import (
	"dev11/internal/domain"
	"dev11/internal/repository"
	"time"
)

var _ EventService = (*eventService)(nil)

// Сервис календаря
type eventService struct {
	repo repository.EventRepository
}

func NewEventService(repo repository.EventRepository) *eventService {
	return &eventService{
		repo: repo,
	}
}

func (s *eventService) Create(event domain.Event) (int, error) {
	return s.repo.Create(event)
}

func (s *eventService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *eventService) EventsForDay(day time.Time) ([]domain.Event, error) {
	return s.repo.EventsForDay(day)
}

func (s *eventService) EventsForMonth(month time.Time) ([]domain.Event, error) {
	return s.repo.EventsForMonth(month)
}

func (s *eventService) EventsForWeek(week time.Time) ([]domain.Event, error) {
	return s.repo.EventsForWeek(week)
}

func (s *eventService) Update(id int, event domain.Event) error {
	return s.repo.Update(id, event)
}
