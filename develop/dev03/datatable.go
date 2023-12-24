package main

import (
	"fmt"
	"sort"
	"strings"
)

// type Sort interface {
//      Len() int
//      Less (i, j) bool
//      Swap(i, j int)
// }

// Структура, содержащая данные и стратегию их сортировки
// Удовлетворяет интерфейсу Sort
type DataTable struct {
	data         [][]string
	sortStrategy SortStrategy
}


func NewDataTable(data [][]string) *DataTable {
	t := &DataTable{
		data: data,
	}

	// По дефолту сортировка по строкам, по 1 столбцу
	t.sortStrategy = StringSortStrategy{t: t}

	return t
}

// Установка стратегии сортировки
func (t *DataTable) SetSortStrategy(s SortStrategy) {
	t.sortStrategy = s
}

func (t DataTable) Data() [][]string {
	return t.data
}

func (t DataTable) SortStrategy() SortStrategy {
	return t.sortStrategy
}

func (t DataTable) Len() int {
	return len(t.data)
}

func (t DataTable) Less(i, j int) bool {
	return t.sortStrategy.Less(&t, i, j)
}

func (t DataTable) Swap(i, j int) {
	t.data[i], t.data[j] = t.data[j], t.data[i]
}

func (t DataTable) String() string {
	strs := make([]string, t.Len())
	for i := range t.data {
		strs[i] = strings.Join(t.data[i], " ")
	}
	return strings.Join(strs, "\n")
}

func (t *DataTable) Sort() {
	sort.Sort(t)
}

func (t *DataTable) ReverseSort() {
	sort.Sort(sort.Reverse(t))
}

func (t *DataTable) Unique() {
	m := make(map[string]bool)
    uniques := make([][]string, 0)

    for _, row := range t.data {
		rowString := fmt.Sprint(row)
        if _, ok := m[rowString]; !ok {
            m[rowString] = true
            uniques = append(uniques, row)
        }
    }

	t.data = uniques
}

func (t DataTable) IsSorted() bool {
	return sort.IsSorted(t)
}
