package main

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	defaultFields    = ""
	defaultDelimeter = ","
	defaultSeparated = false

	fieldsSep      = ","
	fieldsRangeSep = "-"
)

// Функция-аналог cut
func runCut(s, fields, delim string, separated bool) (string, error) {
	// Если включен вывод только с разделителями и в строке их нет, то ничего не выводим
	if separated && !strings.Contains(s, delim) {
		return "", nil
	}

	// Если номера полей не указаны, то выводим всю строку
	if fields == "" {
		return s, nil
	}

	// Разделяем строку по разделителю
	vals := strings.Split(s, delim)
	// Парсим номера полей
	selFields, err := parseFields(fields, fieldsSep, fieldsRangeSep)
	if err != nil {
		return "", err
	}

	// Если номера полей выходят за границы, то ошибка
	if slices.Min(selFields) < 1 || slices.Max(selFields) > len(vals) {
		return "", fmt.Errorf("some fields out of bounds: min = %d, max = %d", 1, len(vals))
	}

	// Записываем в результат выбранные по номерам полей строки
	r := make([]string, 0)
	for _, fieldNum := range selFields {
		r = append(r, vals[fieldNum-1])
	}

	return strings.Join(r, delim), nil
}

// Парсинг полей параметра -f
// Могут быть как единичные значения, так и диапазоны
// Порядок перечисления полей имеет значение
// Примеры: 1 ; 1,2,5 ; 3-5,1 ; 5-5
func parseFields(fields, sep, rangeSep string) ([]int, error) {
	// Если ни один из разделительей не является пробелом
	// То чистим лишние
	if fieldsSep != " " && rangeSep != " " {
		fields = strings.ReplaceAll(fields, " ", "")
	}

	// Разделяем поля
	strs := strings.Split(fields, sep)
	fs := make([]int, 0, len(strs))

	for _, s := range strs {
		// Если попался диапазон - созадем несколько полей
		if strings.Count(s, rangeSep) != 0 {
			rangeFields, err := parseFieldRange(s, rangeSep)
			if err != nil {
				return nil, err
			}

			fs = append(fs, rangeFields...)
		// Иначе просто парсим номер поля
		} else {
			n, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}

			fs = append(fs, n)
		}
	}

	return fs, nil

}

// Парсинг номеров полей в виде диапазонов, например, 1-7 или 2-2
func parseFieldRange(r, rangeSep string) ([]int, error) {
	ns := strings.SplitN(r, rangeSep, 2)
	if len(ns) != 2 {
		return nil, fmt.Errorf("wrong range syntax: no sep in %s", r)
	}

	start, err := strconv.Atoi(ns[0])
	if err != nil {
		return nil, err
	}

	end, err := strconv.Atoi(ns[1])
	if err != nil {
		return nil, err
	}

	if start > end {
		return nil, fmt.Errorf("wrong range: start = %d > end = %d", start, end)
	}

	// Диапазон с включением
	// "1-5" = 1 2 3 4 5,
	// "1-1" = 1
	rng := make([]int, 0, end-start+1)
	for i := start; i <= end; i++ {
		rng = append(rng, i)
	}

	return rng, nil
}

func main() {
	var (
		fields, delimiter string
		separated         bool
	)

	flag.StringVar(&fields, "f", defaultFields, "номера полей [1,2,5]")
	flag.StringVar(&delimiter, "d", defaultDelimeter, "разделитель [,]")
	flag.BoolVar(&separated, "s", defaultSeparated, "только строки с разделителем [true]")
	flag.Parse()

	in := os.Stdin
	out := os.Stdout
	outErr := os.Stderr
	sc := bufio.NewScanner(in)

	for sc.Scan() {
		r, err := runCut(sc.Text(), fields, delimiter, separated)
		if err != nil {
			fmt.Fprint(outErr, err)
		}

		if r != "" {
			fmt.Fprintln(out, r)
		}
	}
}
