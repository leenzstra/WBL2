package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
	Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

/*
	Паттерн State - поведенческий паттерн.
	Позволяет объектам менять поведение в зависимости от своего состояния.
	Применяется когда в объекте есть различные варианты поведения.

	+ разделение отвественности
	+ уменьшение кол-ва ветвлений
	+ расширяемость

	- рост числа классов
	- рост сложности отслеживания состояний в больших системах
	- тяжело обмениться данными межлу состояниями

	Примеры:
	* UI (нажатие, зажатие, выделение)
	* имитация
*/

// Интерфейс состояния воды
type State interface {
	Drink()
	// Нагревание воды
	Heat(*Water)
	// Охлаждение воды
	Freeze(*Water)
}

// Жидкое состояние воды
type LiquidState struct{}

func NewLiquidState() LiquidState {
	return LiquidState{}
}

func (s LiquidState) Heat(w *Water) {
	w.SetState(NewGasState())
}

func (s LiquidState) Freeze(w *Water) {
	w.SetState(NewSolidState())
}

func (s LiquidState) Drink() {
	fmt.Println("Пью нормальную воду")
}

// Твердое состояние воды
type SolidState struct{}

func NewSolidState() SolidState {
	return SolidState{}
}

func (s SolidState) Heat(w *Water) {
	w.SetState(NewLiquidState())
}

func (s SolidState) Freeze(w *Water) {
	fmt.Println("Куда еще морозить")
}

func (s SolidState) Drink() {
	fmt.Println("Кушаю лед")
}

// Газообразное состояние воды
type GasState struct{}

func NewGasState() GasState {
	return GasState{}
}

func (s GasState) Heat(w *Water) {
	fmt.Println("Куда еще нагревать")
}

func (s GasState) Freeze(w *Water) {
	w.SetState(NewLiquidState())
}

func (s GasState) Drink() {
	fmt.Println("Вдыхаю пар")
}

// Структура Вода, использующая состояния
type Water struct {
	state State
}

func NewWater(init State) Water {
	return Water{
		state: init,
	}
}

func (w *Water) SetState(s State) {
	w.state = s
}

func (w Water) State() State {
	return w.state
}

func (w *Water) Heat() {
	w.state.Heat(w)
}

func (w *Water) Freeze() {
	w.state.Freeze(w)
}

func (w *Water) Drink() {
	w.state.Drink()
}

// Пример использования
func StateUsage() {
	water := NewWater(NewLiquidState())
	water.Drink() 	// Пью нормальную воду

	water.Heat()
	water.Drink() 	// Вдыхаю пар

	water.Heat()  	// Куда еще нагревать
	water.Drink() 	// Вдыхаю пар

	water.Freeze()
	water.Drink() 	// Пью нормальную воду

	water.Freeze()
	water.Drink() 	// Кушаю лед

	water.Freeze() 	// Куда еще морозить
	water.Drink()  	// Кушаю лед
}
