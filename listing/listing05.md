Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Интерфейс равен nil, только если тип и значение равны nil.
test() возвращает интерфейс, в котором указатель на данные = nil, но тип будет customError, а не error.

Испрвить можно изменением возвращаемого типа на error
или обявлением var err *customError

```
