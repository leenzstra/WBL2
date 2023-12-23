package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
	Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
	Паттерн Chain of responsibility - поведенческий паттерн.
	Позволяем выполнять операцию с помощью цепочки обработчиков.
	Обработчики решают, передать объект дальше или закончить обработку.
	Динамическая цепочка объектов-обработчиков.

	+ разделение отвественности
	+ гибкость настройки
	+ масштабируемость

	- возможность необработки запроса
	= повышение сложности

	Примеры:
	* http handlers (например, аутентификация)
	* UI
*/

// Предмет, который пройдет через ряд операций
// перед тем как оптправится к покупателю
type Item struct {
	name  string
	price int
}

// Обертка для шагов выполнения операций
type ItemTask struct {
	item     Item
	bought   bool
	packaged bool
	sent     bool
}

// Интерфейс обработчика
type handler interface {
	Execute(*ItemTask)
	SetNext(handler)
}

// Базовая реализвция обработчика
type handlerBase struct {
	next handler
}

// Установка следющего обработчика в цепочке
func (h *handlerBase) SetNext(next handler) {
	h.next = next
}

// Вызов следующего обработчика в цепочке, если он есть
func (h *handlerBase) Execute(task *ItemTask) {
	if h.next != nil {
		h.next.Execute(task)
	}
}

// Сервис покупки предмета
type PayService struct {
	handlerBase
}

func (ps *PayService) Execute(task *ItemTask) {
	if !task.bought {
		fmt.Println("Предмет куплен:", task.item.name)
		task.bought = true
	}
	ps.handlerBase.Execute(task)
}

func (ps *PayService) SetNext(next handler) {
	ps.next = next
}

// Сервис упаковки предмета
type PackagingService struct {
	handlerBase
}

func (ps *PackagingService) Execute(task *ItemTask) {
	if !task.packaged {
		fmt.Println("Предмет упакован:", task.item.name)
		task.packaged = true
	}
	ps.handlerBase.Execute(task)
}

func (ps *PackagingService) SetNext(next handler) {
	ps.next = next
}

// Сервис отправки предмета
type DeliveryService struct {
	handlerBase
}

func (ds *DeliveryService) Execute(task *ItemTask) {
	if !task.sent {
		fmt.Println("Предмет отправлен:", task.item.name)
		task.sent = true
	}
	ds.handlerBase.Execute(task)
}

func (ds *DeliveryService) SetNext(next handler) {
	ds.next = next
}

// Пример использования
func ChainUsage() {
	item := Item{"ксяоми", 7999}
	task := &ItemTask{item: item}

	ps := &PayService{}
	pkgs := &PackagingService{}
	ds := &DeliveryService{}

	ps.SetNext(pkgs)
	pkgs.SetNext(ds)

	ps.Execute(task)
}
