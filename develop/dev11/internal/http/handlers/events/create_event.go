package events

import (
	"dev11/internal/http/extract"
	"dev11/internal/http/response"
	"fmt"
	"net/http"
)

// Обаботчик создания нового события
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	// Должен быть метод POST
	if r.Method != http.MethodPost {
		response.SendJSONErr(w, r, fmt.Errorf("method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	// Достаем параметры запроса (поля Event)
	event, err := extract.CreateEventFormData(r)
	if err != nil {
		response.SendJSONErr(w, r, err, http.StatusBadRequest)
		return
	}

	// Добавляем событие в базу
	id, err := h.service.Create(event)
	if err != nil {
		response.SendJSONErr(w, r, err, http.StatusServiceUnavailable)
		return
	}

	// Возвращаем id созданного события
	response.SendJSON(w, r, map[string]int{"id": id}, http.StatusCreated)
}
