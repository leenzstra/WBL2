package main

import (
	"fmt"
	"sync"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

// Функция объединения нескольких done каналов
// Закрывает первый выполнившийся done
func or(channels ...<-chan interface{}) <-chan interface{} {
	// Объединяющий done канал
	out := make(chan interface{})
	// Если каналов не было передано, то сразу его закрываем
	if len(channels) == 0 {
		close(out)
		return out
	}

	// Объект для закрытия канала единожды
	var o sync.Once

	// Читаем все каналы
	for _, ch := range channels {
		go func(ch <-chan interface{}) {
			select {
			// Если пришел сигнал от одного из каналов, за закрываем основной
			case <-ch:
				o.Do(func() {
					close(out)
				})
			// Если основной закрылся, то выход из селекта
			case <-out:
				break
			}
		}(ch)
	}

	return out
}

// Закрытие канала после after
func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func main() {
	s := time.Now()

	<-or(
		sig(2*time.Hour),
		sig(1*time.Second),
		sig(2*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	e := time.Since(s).Milliseconds()

	fmt.Printf("elapsed: %d ms\n", e)
}
