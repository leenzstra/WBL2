package events

import (
	"dev11/internal/domain"
	"dev11/internal/http/response"
	"dev11/pkg/dateutils"
	"fmt"
	"net/http"
	"time"
)

func (h *handler) EventsForMonth(w http.ResponseWriter, r *http.Request) {
	// Должен быть метод GET
	if r.Method != http.MethodGet {
		response.SendJSONErr(w, r, fmt.Errorf("method not allowed"), http.StatusMethodNotAllowed)
		return
	}
	
	// Заданный день
	dateStr := r.URL.Query().Get("date")
	date, err := time.Parse(dateutils.DefaultLayout, dateStr)
	if err != nil {
		response.SendJSONErr(w, r, err, http.StatusBadRequest)
		return
	}

	// Получение событий
	events, err := h.service.EventsForMonth(date)
	if err != nil {
		response.SendJSONErr(w, r, err, http.StatusServiceUnavailable)
		return
	}
	// Конвератция в ответы сервера
	eventsResp := make([]domain.EventResponse, len(events))
	for i, e := range events {
		eventsResp[i] = domain.NewEventResponse(e)
	}

	response.SendJSON(w, r, eventsResp, http.StatusOK)
}
