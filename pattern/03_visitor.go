package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
	Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
	Паттерн Visitor - поведенческий паттерн проектирования.
	Позволяет отделить алгоритм от объекта, над которым он работает.
	Полезен, когда есть сложная структура и мы хотим выполнить над ней некоторые операции, не сильно меняя саму структуру.

	+ не изменяем существующие классы
	+ разделение ответственности

	- тяжело реализовать с приватными полями
	- усложнение кода

	Примеры:
	* работа с файловой системой
*/

// Интерфейс посетителя
type Visitor interface {
	VisitTree(Tree)
}

// Реализация посетителя для рассчета кол-во древесины с дерева
type WoodAmountCalculator struct{}

// Посещение дерева t
func (c *WoodAmountCalculator) VisitTree(t Tree) {
	fmt.Println("Кол-во древесины:", t.Params().Height*t.Params().Width*t.Params().Width)
}

// Интерфейс дерева
type Tree interface {
	Params() TreeParams
	Accept(Visitor)
}

// Параметры дерева
type TreeParams struct {
	Height  int
	Width   int
	Wet     float64
	Chopped bool
}

// Основа всех деревьев, база
// реализует интерфейс Tree
type TreeBase struct {
	params TreeParams
}

func (tb *TreeBase) Params() TreeParams {
	return tb.params
}

// Посещение дерева любым объектом, реализующим интерфейс Visitor
func (tb *TreeBase) Accept(v Visitor) {
	v.VisitTree(tb)
}

// Дуб, встраивает TreeBase
type Oak struct {
	TreeBase
}

// Березка, встраивает TreeBase
type Birch struct {
	TreeBase
}

// Пример использования
func VisitorUsage() {
	// Создаем объекты деревьев
	oakParams := TreeParams{2, 2, 0.1, false}
	oak := Oak{TreeBase{oakParams}}

	birchParams := TreeParams{4, 1, 0.5, false}
	birch := Birch{TreeBase{birchParams}}

	// Создаем посетителя
	calculator := &WoodAmountCalculator{}

	// Принимаем с распростертыми объятиями
	oak.Accept(calculator)
	birch.Accept(calculator)
}
