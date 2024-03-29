Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
nil
false

Foo() возвращает значение nil типа *os.PathError. 
Любое значение == nil только тогда, когда и значение и тип являются nil (интерфейс содержит тип и указатель на данные). 
Поэтому сравнение err == nil будет false, т.к. *os.PathError != nil.

Интерфейс - это структура (iface), которая содержит метаданные (itable) и указатель на данные.

Пустой интерфейс - интерфейс, не имеющий методов, поэтому его реализуют все типы данных.
Для описания пустого интерфейса используется структура eface, содержащая только тип и указатель на данные.

```
