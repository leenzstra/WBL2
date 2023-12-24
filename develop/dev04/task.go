package main

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"fmt"
	"slices"
	"strings"
)

type AnagramFinder struct {
	words []string
}

func NewAnagramFinder(words []string) *AnagramFinder {
	return &AnagramFinder{
		words: words,
	}
}

// Подготовка данных:
// Примведение в lowercase и удаление дубликатов
func (a *AnagramFinder) prepare() []string {
	set := make(map[string]struct{})
	newWords := make([]string, 0)

	for _, w := range a.words {
		wlc := strings.ToLower(w)
		// Удаление дубликатов
		if _, exists := set[wlc]; !exists {
			set[wlc] = struct{}{}
			newWords = append(newWords, wlc)
		}
	}

	return newWords
}

// Проверка являются ли слвова анаграммой
func (a *AnagramFinder) isAnagram(w1, w2 string) bool {
	if len(w1) != len(w2) {
		return false
	}

	// Подсчет символов в каждом слове
	w1Runes := make(map[rune]int)
	w2Runes := make(map[rune]int)

	for _, r := range w1 {
		w1Runes[r]++
	}

	for _, r := range w2 {
		w2Runes[r]++
	}

	// Проверяем, что кол-во уникальных символов совпадает
	if len(w1Runes) != len(w2Runes) {
		return false
	}

	// Сравниваем кол-во символов
	for r1, count1 := range w1Runes {
		if count2, ok := w2Runes[r1]; !ok || count1 != count2 {
			return false
		}
	}

	return true
}

// Поиск групп анаграмм
func (a *AnagramFinder) FindGroups() map[string][]string {
	// Подготовка слов
	prepWords := a.prepare()
	// Выделение памяти под группы
	groups := make(map[string][]string)

	// Каждое слово сравниваем с ключами групп
	// Если ключ и слово = анаграмма, то добавляем в массив группы
	// Если анаграммы не нашлось, то добавляем новую группу
	for _, word := range prepWords {
		foundAnagram := false

		for k := range groups {
			if a.isAnagram(k, word) {
				groups[k] = append(groups[k], word)
				foundAnagram = true
				break
			}
		}	
		if !foundAnagram {
			groups[word] = append(groups[word], word)
		}
	}

	// Удаление множетств из 1 элемента и сортировка остальных
	for k, v := range groups {
		if len(v) == 1 {
			delete(groups, k)
		} else {
			slices.Sort(v)
		}
	}

	return groups
}

func main() {
	data := []string{"пятак", "слиток","ПЯТКА", "пятка", "тяпка", "листок", "СтоЛИК", "листок","тилокс", "игорь","гигорь","гогирь"}
	finder := NewAnagramFinder(data)
	for k, v := range finder.FindGroups() {
		fmt.Printf("%s:\t %v\n", k, v)
	}
}
