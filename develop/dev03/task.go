package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func splitRow(s, sep string) []string {
	return strings.Split(s, sep)
}

// Читаем исходные данные из файла
func readInput(file string) [][]string {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	data := make([][]string, 0)

	for sc.Scan() {
		data = append(data, splitRow(sc.Text(), " "))
	}

	return data
}

func main() {
	var (
		inputFile                                        string
		sortColumn                                       int
		numeric, reverse, unique, monthSort              bool
		ignoreTrailingSpaces, checkSorted, numericSuffix bool
	)

	flag.StringVar(&inputFile, "i", "input.txt", "input file")

	flag.IntVar(&sortColumn, "k", 0, "sort column index")
	flag.BoolVar(&numeric, "n", false, "numeric sort")
	flag.BoolVar(&reverse, "r", false, "reverse sort")
	flag.BoolVar(&unique, "u", false, "unique only")

	flag.BoolVar(&monthSort, "M", false, "month name sort")
	flag.BoolVar(&ignoreTrailingSpaces, "b", false, "ignore trailing spaces")
	flag.BoolVar(&checkSorted, "c", false, "check is sorted")
	flag.BoolVar(&numericSuffix, "h", false, "number + suffix sort")

	// Парсим аргументы
	flag.Parse()

	// Читаем данные их файла
	data := readInput(inputFile)

	// Создаем объект таблицы с данными
	t := NewDataTable(data)

	// Выбираем стратегию сортировки
	if numeric {
		t.SetSortStrategy(NumericSortStrategy{t: t, col: sortColumn,
			ignoreTrailingSpaces: ignoreTrailingSpaces})
	} else if monthSort {
		t.SetSortStrategy(MonthSortStrategy{t: t, col: sortColumn,
			ignoreTrailingSpaces: ignoreTrailingSpaces})
	} else if numericSuffix {
		t.SetSortStrategy(NumberSufSortStrategy{t: t, col: sortColumn,
			ignoreTrailingSpaces: ignoreTrailingSpaces})
	} else {
		t.SetSortStrategy(StringSortStrategy{t: t, col: sortColumn,
			ignoreTrailingSpaces: ignoreTrailingSpaces})
	}

	// Сортируем в выбранном порядке
	if reverse {
		t.ReverseSort()
	} else {
		t.Sort()
	}

	// Выбираем только уникальные строки
	if unique {
		t.Unique()
	}

	// Проверка на отсортированность
	if checkSorted {
		fmt.Println("sorted:", t.IsSorted())
	}

	fmt.Println(t)

}
