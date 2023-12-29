package events

import (
	"dev11/internal/http/response"
	"fmt"
	"net/http"
	"strconv"
)

// Обработчик удаления события
func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	// Должен быть метод POST
	if r.Method != http.MethodPost {
		response.SendJSONErr(w, r, fmt.Errorf("method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	// Достаем параметры запроса (id события)
	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.SendJSONErr(w, r, err, http.StatusBadRequest)
		return
	}

	// Удаляем событие из базы
	err = h.service.Delete(id)
	if err != nil {
		response.SendJSONErr(w, r, err, http.StatusServiceUnavailable)
		return
	}

	response.SendJSON(w, r, true, http.StatusOK)
}
