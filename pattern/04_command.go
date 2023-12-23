package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
	Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
	Паттерн Command - поведенческий паттерн.
	Запросы представляются объектами, которые можно передать и выполнить позже.
	Применяется для возможности отложенного выполнения операций, отмены, унификации

	+ разделение отвественности
	+ отложенный запуск или отмена

	- усложнение структуры кода

	Примеры:
	* golang exec.Command()
	* golang http.NewRequest()
*/

// OrderCommand представляет интерфейс команды
type OrderCommand interface {
	Execute()
}

// BuyStockCommand представляет команду покупки акций
type BuyStockCommand struct {
	stock *Stock
	price float64
}

func (bsc *BuyStockCommand) Execute() {
	bsc.stock.Buy(bsc.price)
}

// SellStockCommand представляет команду продажи акций
type SellStockCommand struct {
	stock *Stock
	price float64
}

func (ssc *SellStockCommand) Execute() {
	ssc.stock.Sell(ssc.price)
}

// Stock представляет объект акций
type Stock struct {
	name     string
	quantity int
}

func (s *Stock) Buy(price float64) {
	fmt.Printf("Акция: %s, Кол-во: %d куплено\n", s.name, s.quantity)
}

func (s *Stock) Sell(price float64) {
	fmt.Printf("Акция: %s, Кол-во: %d продано\n", s.name, s.quantity)
}

// Broker представляет вызывающий объект
type Broker struct {
	orders []OrderCommand
}

func (b *Broker) TakeOrder(order OrderCommand) {
	b.orders = append(b.orders, order)
}

func (b *Broker) PlaceOrders() {
	for _, order := range b.orders {
		order.Execute()
	}
	b.orders = nil
}

// Пример использования
func CommandUsage() {
	stock := &Stock{name: "AAPL", quantity: 10}

	buyStockOrder := &BuyStockCommand{stock: stock, price: 100}
	sellStockOrder := &SellStockCommand{stock: stock, price: 110}

	broker := &Broker{}
	broker.TakeOrder(buyStockOrder)
	broker.TakeOrder(sellStockOrder)

	broker.PlaceOrders()
}
