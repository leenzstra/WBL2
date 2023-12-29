package events

import (
	"dev11/internal/http/middleware"
	"dev11/internal/service"
	"net/http"
)

type handler struct {
	service service.EventService
}

func NewEventsHandler(service service.EventService) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) SetupRoutes() http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("/create_event", middleware.LogMiddleware(h.Create))
	r.HandleFunc("/update_event", middleware.LogMiddleware(h.Update))
	r.HandleFunc("/delete_event", middleware.LogMiddleware(h.Delete))
	r.HandleFunc("/events_for_day", middleware.LogMiddleware(h.EventsForDay))
	r.HandleFunc("/events_for_week", middleware.LogMiddleware(h.EventsForWeek))
	r.HandleFunc("/events_for_month", middleware.LogMiddleware(h.EventsForMonth))

	return r
}
