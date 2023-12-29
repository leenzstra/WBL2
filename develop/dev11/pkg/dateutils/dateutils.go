package dateutils

import "time"

const DefaultLayout = "2006-01-02"

// Возвращает начало текущего месяца и начало следующего месяца по заданному дню
func MonthRange(day time.Time) (time.Time, time.Time) {
	y, m, _ := day.Date()
	startMonth := time.Date(y, m, 1, 0, 0, 0, 0, day.Location())
	nextStartMonth := startMonth.AddDate(0, 1, 0)

	return startMonth, nextStartMonth
}

// Возвращает начало текущей недели и начало следующей недели по заданному дню
// Начало - понедельник
func WeekRange(day time.Time) (time.Time, time.Time) {
	var startWeek, nextStartWeek time.Time

	if day.Weekday() == time.Sunday {
		startWeek = day.AddDate(0, 0, -6)
	} else {
		startWeek = day.AddDate(0, 0, int(time.Monday - day.Weekday()))
	}

	nextStartWeek = startWeek.AddDate(0, 0, 7)

	return startWeek, nextStartWeek
}