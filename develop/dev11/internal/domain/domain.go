package domain

import (
	"dev11/pkg/dateutils"
	"time"
)

// Структура события
type Event struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}

// Структура ответа от сервера
type EventResponse struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Date        string `json:"date"`
	Description string `json:"description"`
}

func NewEventResponse(e Event) EventResponse {
    return EventResponse{
    	Id:          e.Id,
    	Title:       e.Title,
    	Date:        e.Date.Format(dateutils.DefaultLayout),
    	Description: e.Description,
    }
}
