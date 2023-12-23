package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
	Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

/*
	Паттерн Strategy - поведенческий паттерн.
	Определяет семейство алгоритмов и обсеспечивает их взаимозаменяемость.
	Это позволяет выбирать алгоритм путём определения соответствующего класса.

	Зачем:
	* Программа должна обеспечивать различные варианты алгоритма или поведения
	* Нужно изменять поведение каждого экземпляра класса
	* Необходимо изменять поведение объектов на стадии выполнения

	+ изменение поведения во время выполнения
	+ инкапсуляция
	+ упрощение модификации 

	- усложнение за счет доп. классов
	- незнание особенностей стратегий со стороны пользователя

	Примеры:
	* sort
	* trading bots
	* оплата
*/

// Интерфейс стратегии
type EatStrategy interface {
	Eat()
}

// Стратегия еды ложкой
type SpoonEatStrategy struct{
	chunkSize int
}

func (s SpoonEatStrategy) Eat() {
	fmt.Printf("Ем ложкой %d за раз\n", s.chunkSize)
}

// Стратугия еды вилкой
type ForkEatStrategy struct{
	chunkSize int
}

func (s ForkEatStrategy) Eat() {
	fmt.Printf("Ем вилкой %d за раз\n", s.chunkSize)
}

type Cake struct {
	strategy EatStrategy
}

func NewCake(s EatStrategy) *Cake {
	return &Cake{
		strategy: s,
	}
}

func (c *Cake) Eat() {
	c.strategy.Eat()
}

// Пример использования
func StrategyUsage() {
	forkStrat := ForkEatStrategy{1}
	spoonStrat := SpoonEatStrategy{2}

	cake := NewCake(forkStrat)
	cake.Eat()

	cake.strategy = spoonStrat
	cake.Eat()
}
