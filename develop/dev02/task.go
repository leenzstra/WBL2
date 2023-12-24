package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func StringUnpack(s string) (string, error) {
	// Билдер результата
	builder := strings.Builder{}
	// Билдер числа, кол-ва дубликатов символов
	numBuilder := strings.Builder{}
	runes := []rune(s)

	// Если строка пустая, то ее и возвращаем
	if s == "" {
		return "", nil
	}

	// Если первая руна = число, то ошибка
	if unicode.IsDigit(runes[0]) {
		return "", fmt.Errorf("wrong string starts with: %s", string(runes[0]))
	}

	// Проходимся по каждой руне
	for i := 0; i < len(runes); i++ {
		// Если руна - не число и не экскейп, добавляем его в результат
		if unicode.IsLetter(runes[i]) {
			builder.WriteRune(runes[i])
		// Если руна эскейп, то добавляем следующую руну, если возможно, 
		// и двигаемся на 1 шаг вперед
		} else if runes[i] == '\\' {
			if i+1 < len(runes) {
				builder.WriteRune(runes[i+1])
			}
			i++
		// В остальных случаях (число)
		} else {
			// Собираем число целиком
			for j := i; j < len(runes) && unicode.IsDigit(runes[j]); j++ {
				numBuilder.WriteRune(runes[j])
			}
			
			// Парсим его в int
			n, err := strconv.Atoi(numBuilder.String())
			if err != nil {
				return "", fmt.Errorf("parse int error: %s", err.Error())
			}

			// Дублируем предыдущую руну N раз
			for j := 0; j < n-1; j++ {
				builder.WriteRune(runes[i-1])
			}
			
			// Сдвигаемся на длину числа вперед
			i += len(numBuilder.String()) - 1
			// Обнуляем билдер кол-ва дубликатов
			numBuilder.Reset()
		}
	}

	return builder.String(), nil
}