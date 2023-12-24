package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

// Примеры вызова
// go run . 
// go run . -ntp <addr>

const (
	// Дефолтный сервер
	defaultNtpAddr = "0.beevik-ntp.pool.ntp.org"
)

func GetTime(addr string) (time.Time, error) {
	return ntp.Time(addr)
}

func main() {
	// Получаем адрес ntp сервера из параметров
	var addr string
	flag.StringVar(&addr, "ntp", defaultNtpAddr, "ntp server address")
	flag.Parse()

	// Получение времени с сервера
	t, err := GetTime(addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "get time error: %v\n", err)
		os.Exit(1)
	}

	// Вывод времени
	fmt.Println(t.Format(time.RFC3339))
}