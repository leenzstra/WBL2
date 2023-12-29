package events

import (
	"dev11/internal/http/extract"
	"dev11/internal/http/response"
	"fmt"
	"net/http"
)

// Обработчик обновления события
func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	// Должен быть метод POST
	if r.Method != http.MethodPost {
		response.SendJSONErr(w, r, fmt.Errorf("method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	// Достаем параметры запроса (поля Event)
	event, err := extract.UpdateEventFormData(r)
	if err != nil {
		response.SendJSONErr(w, r, err, http.StatusBadRequest)
		return
	}

	// Обновляем событие в базе
	err = h.service.Update(event.Id, event)
	if err != nil {
		response.SendJSONErr(w, r, err, http.StatusServiceUnavailable)
		return
	}

	response.SendJSON(w, r, true, http.StatusOK)
}
