Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
Числа от 1 до 8 в случайном порядке, затем много 0.

Чтение из закрытого канала будет возвращать zeroValue типа канала (int) и false вторым значением (канал закрыт).
Также канал C не закрывается, из-за чего цикл range бесконечно продолжает выводить нули из A и B.

Исправить можно путем закрытия канала С в merge и отслеживанием закрытия каналов A и B.
...



```
