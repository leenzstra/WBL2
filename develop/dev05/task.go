package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения !
-B - "before" печатать +N строк до совпадения !
-C - "context" (A+B) печатать ±N строк вокруг совпадения !
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр) !
-v - "invert" (вместо совпадения, исключать) !
-F - "fixed", точное совпадение со строкой, не паттерн !
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Читаем исходные данные из файла
func readInput(file string) []string {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	data := make([]string, 0)

	for sc.Scan() {
		data = append(data, sc.Text())
	}

	return data
}

// Поиск совпадений в строках (ignoreCase, invert, fixed)
func findMatches(lines []string, pattern string, ignoreCase, invert, fixed bool) (matches []int, err error) {
	var re *regexp.Regexp

	// Меняем регулярное выражение в зависимости от ignoreCase
	if ignoreCase {
		re, err = regexp.Compile(strings.ToLower(pattern))
	} else {
		re, err = regexp.Compile(pattern)
	}
	if err != nil {
		return nil, err
	}

	for i, line := range lines {
		// Ищем совпадение в зависмости от параметров
		idxMatched := false
		// Если fixed, то ищем полное совпадение строки
		if fixed {
			idxMatched = line == pattern
		// Если ignoreCase, то ищем совпадение строки и паттерна в нижнем регистре
		} else if ignoreCase {
			idxMatched = re.Match([]byte(strings.ToLower(line)))
		// Если ни то, ни другое, то ищем совпадение строки и паттерна в изначальных формах
		} else {
			idxMatched = re.Match([]byte(line))
		}

		// Инвертируем совпадение
		if invert {
			idxMatched = !idxMatched
		}

		// Добавляем совпадение (индекс строки!!!)
		if idxMatched {
			matches = append(matches, i)
		}
	}

	return
}

// Добавление окружеющих строк (before, after, contexr)
func addOffset(matches []int, maxLines, before, after, context int) (newMatches []int) {
	// Если указаны before, after и context
	// то выбираем наибольшнее смещение
	if context > before {
		before = context

	}
	if context > after {
		after = context
	}

	for _, match := range matches {
		// Добавить строки ДО
		for j := clamp(match-before,0,match); j != match; j++ {
			newMatches = append(newMatches, j)
		}
		// добавить совпадение
		newMatches = append(newMatches, match)
		// добавить строки ПОСЛЕ
		for j, a := match + 1, 0; j < maxLines && a < after; j, a = j+1, a+1 {
			newMatches = append(newMatches, j)
		}
	}

	return
}

// вывод совпадений (count, lineNum)
func output(writer io.Writer, matches []int, lines []string, count, lineNum bool) {
	for _, matchIdx := range matches {
		// Вывод с номером строки в изначальном файле
		if lineNum {
			fmt.Fprintf(writer, "%d. %s\n", matchIdx+1, lines[matchIdx])
		// Без номера
		} else {
			fmt.Fprintf(writer, "%s\n", lines[matchIdx])
		}
	}

	// В конце вывод кол-ва строк
	if count {
		fmt.Fprint(writer,len(matches))
	}
}

func run(lines []string, pattern string, ignoreCase, invert, fixed, count, lineNum bool, before, after, context int, out io.Writer) {
	// Ищем совпадения
	matches, err := findMatches(lines, pattern, ignoreCase, invert, fixed)
	if err != nil {
		log.Fatal(err)
	}

	// Добавляем оффсеты
	matchesWithOffset := addOffset(matches, len(lines), before, after, context)

	// Выводим куда удобно
	output(out, matchesWithOffset, lines, count, lineNum)
}

func main() {
	var (
		after, before, context                    int
		count, ignoreCase, invert, fixed, lineNum bool
	)

	flag.IntVar(&after, "A", 0, "печатать +N строк после совпадения")
	flag.IntVar(&before, "B", 0, "печатать +N строк до совпадения")
	flag.IntVar(&context, "C", 0, "печатать ±N строк вокруг совпадения")
	flag.BoolVar(&count, "c", false, "количество строк")
	flag.BoolVar(&ignoreCase, "i", false, "игнорировать регистр")
	flag.BoolVar(&invert, "v", false, "вместо совпадения, исключать")
	flag.BoolVar(&fixed, "F", false, "точное совпадение со строкой")
	flag.BoolVar(&lineNum, "n", false, "печатать номер строки")

	flag.Parse()

	// Проверка кол-ва аргументов
	if flag.NArg() < 2 {
		fmt.Println("go run . [опции] <шаблон> <файл>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Паттерн
	pattern := flag.Arg(0)
	// Файл со строками
	inputFile := flag.Arg(1)

	// Читаем троки из файла
	lines := readInput(inputFile)

	run(lines, pattern, ignoreCase, invert, fixed, count, lineNum, before, after, context, os.Stdout)
}
