package main

import (
	"strconv"
	"strings"
	"time"
	"unicode"
)

// Стратегии сортировки данных

type SortStrategy interface {
	Less(t *DataTable, i, j int) bool
}

// Каждая стратегия содержит указатель на структуру данных
// Колонку, по которой сортируем
// и игнорирование пробельных символов

// Числовая сортировка
type NumericSortStrategy struct {
	t                    *DataTable
	col                  int
	ignoreTrailingSpaces bool
}

func (s NumericSortStrategy) Less(t *DataTable, i, j int) bool {
	n1 := s.t.data[i][s.col]
	n2 := s.t.data[j][s.col]

	if s.ignoreTrailingSpaces {
		n1 = strings.TrimRight(n1, " ")
		n2 = strings.TrimRight(n2, " ")
	}

	i1, err := strconv.Atoi(n1)
	if err != nil {
		i1 = 0
	}
	i2, err := strconv.Atoi(n2)
	if err != nil {
		i2 = 0
	}

	return i1 < i2
}

// Строковая сортировка
type StringSortStrategy struct {
	t                    *DataTable
	col                  int
	ignoreTrailingSpaces bool
}

func (s StringSortStrategy) Less(t *DataTable, i, j int) bool {
	n1 := t.data[i][s.col]
	n2 := t.data[j][s.col]

	if s.ignoreTrailingSpaces {
		n1 = strings.TrimRight(n1, " ")
		n2 = strings.TrimRight(n2, " ")
	}

	return n1 < n2
}

// Сортировка по названию месяца
type MonthSortStrategy struct {
	t                    *DataTable
	col                  int
	ignoreTrailingSpaces bool
}

func (s MonthSortStrategy) Less(t *DataTable, i, j int) bool {

	n1 := t.data[i][s.col]
	n2 := t.data[j][s.col]

	if s.ignoreTrailingSpaces {
		n1 = strings.TrimRight(n1, " ")
		n2 = strings.TrimRight(n2, " ")
	}
	t1, err := time.Parse("January", n1)
	if err != nil {
		t1 = time.Time{}
	}
	t2, err := time.Parse("January", n2)
	if err != nil {
		t2 = time.Time{}
	}

	return t1.Before(t2)
}

// Сортировка по числу с суффиксом
type NumberSufSortStrategy struct {
	t                    *DataTable
	col                  int
	ignoreTrailingSpaces bool
}

func (s NumberSufSortStrategy) getSuffix(str string) (int, string) {
	num := ""
	suf := ""

	for i := len(str) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(str[i])) || str[i] == '-' {
			num = string(str[i]) + num
		} else {
			suf = string(str[i]) + suf
		}
	}

	n, err := strconv.Atoi(num)
	if err != nil {
		n = 0
	}

	return n, suf
}

func (s NumberSufSortStrategy) Less(t *DataTable, i, j int) bool {

	n1 := t.data[i][s.col]
	n2 := t.data[j][s.col]

	if s.ignoreTrailingSpaces {
		n1 = strings.TrimRight(n1, " ")
		n2 = strings.TrimRight(n2, " ")
	}

	num1, suf1 := s.getSuffix(n1)
	num2, suf2 := s.getSuffix(n2)

	if num1 == num2 {
		return suf1 < suf2
	} 

	return num1 < num2
}
