package repository

import (
	"dev11/internal/domain"
	"dev11/pkg/dateutils"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var _ EventRepository = (*pgEventRepository)(nil)

// Репозиторий календаря, используеющий Postgres как хранилище
type pgEventRepository struct {
	db *sqlx.DB
}

func NewPgEventRepository(db *sqlx.DB) *pgEventRepository {
	return &pgEventRepository{
		db: db,
	}
}

func (r *pgEventRepository) Create(event domain.Event) (int, error) {
	q := `INSERT INTO events (title, date, description) VALUES ($1, $2, $3) RETURNING id`
	row := r.db.QueryRow(q, event.Title, event.Date, event.Description)

	id := 0
	err := row.Scan(&id)

	return id, err
}

func (r *pgEventRepository) Delete(id int) error {
	q := `DELETE FROM events WHERE id = $1`

	res, err := r.db.Exec(q, id)
	if err != nil {
		return err
	}

	deleted, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if deleted == 0 {
		return fmt.Errorf("not exists: id %d", id)
	}

	return nil
}

func (r *pgEventRepository) EventsForDay(day time.Time) ([]domain.Event, error) {
	dayFormatted := day.Format(dateutils.DefaultLayout)

	q := `SELECT * FROM events WHERE "date" = $1`

	events := []domain.Event{}
	err := r.db.Select(&events, q, dayFormatted)

	return events, err
}

func (r *pgEventRepository) EventsForMonth(day time.Time) ([]domain.Event, error) {
	startMonth, nextStartMonth := dateutils.MonthRange(day)

	smFormatted := startMonth.Format(dateutils.DefaultLayout)
	nsmFormatted := nextStartMonth.Format(dateutils.DefaultLayout)

	q := `SELECT * FROM events WHERE "date" >= $1 AND "date" < $2`

	events := []domain.Event{}
	err := r.db.Select(&events, q, smFormatted, nsmFormatted)

	return events, err
}

// Начало недели - понедельник (1)
// Конец недели - воскресенье (0)
func (r *pgEventRepository) EventsForWeek(day time.Time) ([]domain.Event, error) {
	startWeek, nextStartWeek := dateutils.WeekRange(day)

	q := `SELECT * FROM events WHERE "date" >= $1 AND "date" < $2`

	events := []domain.Event{}
	err := r.db.Select(&events, q, startWeek, nextStartWeek)

	return events, err
}

func (r *pgEventRepository) Update(id int, event domain.Event) error {
	q := `UPDATE events SET`
	n := 1
	args := make([]interface{}, 0)
	stmnts := make([]string, 0)

	if event.Title != "" {
		stmnts = append(stmnts, fmt.Sprintf(" title = $%d", n))
		args = append(args, event.Title)
		n++
	}

	if event.Description != "" {
		stmnts = append(stmnts, fmt.Sprintf(" description = $%d", n))
		args = append(args, event.Description)
		n++
	}

	emptyTime := time.Time{}
	if event.Date != emptyTime {
		stmnts = append(stmnts, fmt.Sprintf(" date = $%d", n))
		args = append(args, event.Date)
		n++
	}

	q += strings.Join(stmnts, ",")
	q += fmt.Sprintf(" WHERE id = $%d", n)
	args = append(args, id)

	_, err := r.db.Exec(q, args...)

	return err
}
