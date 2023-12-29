package extract

import (
	"dev11/internal/domain"
	"dev11/pkg/dateutils"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// Извлечение параметров из www-url-form-encoded для создание события
func CreateEventFormData(r *http.Request) (domain.Event, error) {
	event := domain.Event{}

	title := r.FormValue("title")
	if title == "" {
		return event, fmt.Errorf("title not found")
	}

	desc := r.FormValue("description")
	if desc == "" {
		return event, fmt.Errorf("desc not found")
	}

	d := r.FormValue("date")
	date, err := time.Parse(dateutils.DefaultLayout, d)
	if err != nil {
		return event, fmt.Errorf("incorrect date: %s", d)
	}

	return domain.Event{
		Title:       title,
		Date:        date,
		Description: desc,
	}, nil
}

// Извлечение параметров из www-url-form-encoded для обновления события
func UpdateEventFormData(r *http.Request) (domain.Event, error) {
	event := domain.Event{}

	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return event, fmt.Errorf("incorrect id: %d", id)
	}

	title := r.FormValue("title")
	desc := r.FormValue("description")

	var date time.Time
	if d := r.FormValue("date"); d != "" {
		date, err = time.Parse(dateutils.DefaultLayout, d)
		if err != nil {
			return event, fmt.Errorf("incorrect date: %s", d)
		}
	}

	return domain.Event{
		Id:          id,
		Title:       title,
		Date:        date,
		Description: desc,
	}, nil
}
