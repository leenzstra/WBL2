Что выведет программа? Объяснить вывод программы. Рассказать про внутреннее устройство слайсов и что происходит при передачи их в качестве аргументов функции.

```go
package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}
```

Ответ:
```
[3 2 3]

Передача в фунцию слайса идет по значению, s и i - разные переменные, но они указывают на один массив.
- i[0] изменит первый элемент в обоих слайсах. [3 2 3]
- Емкость слайса = 3. При добавление элемента 4 емкости НЕ хватит, поэтому будет выделен новый массив и слайс с емкостью 6. [3 2 3 4]
- i[1] изменит второй элемент только в новом слайсе. [3 5 3 4]
- Емкость слайса = 6. При добавление элемента 6 емкости хватит, поэтому будет добавление в этот же слайс. [3 2 3 4 6]

Слайс состоит из указателя на массив, длины (текущей длины слайса) и емкости (длины внутреннего массива).
В отличие от массива может менять свой размер и динамически выделять память.

Добавление в слайс происходит с помощью функции append(). При недостаточной емкости выделяется больше памяти, 
создается новый слайс, указывающий на новый массив.


```