Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
		
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
2
1

Правило: Отложенные функции могут считывать и присваивать именованные возвращаемые значения возвращаемой функции.

Это правило работает в первом случае.

Во втором случае сначала выполяется return (когда x = 1), а уже потом defer, но возвращемое значение 
она уже не может увеличить, а только локальный x (который станет = 2).

```
