package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные из сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

const (
	defaultTimeout = 10 * time.Second
	defaultHost    = "localhost"
	defaultPort    = 0
)

type Telnet struct {
	addr    string
	timeout time.Duration
	conn    net.Conn
}

func NewTelnet(addr string, timeout time.Duration) *Telnet {
	return &Telnet{
		addr:    addr,
		timeout: timeout,
	}
}

// Открытие подключения к хосту
func (t *Telnet) Connect() error {
	conn, err := net.DialTimeout("tcp", t.addr, t.timeout)
	if err != nil {
		return err
	}
	t.conn = conn
	return nil
}

func (t *Telnet) Disconnect() {
	t.conn.Close()
}

// Чтение из from, запись в to с возможностью отмены
// Чтение по linebreak'y
func (t *Telnet) rw(ctx context.Context, from io.Reader, to io.Writer) error {
	r := bufio.NewReader(from)

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			d, err := r.ReadString('\n')
			if err != nil {
				return err
			}

			if _, err := fmt.Fprint(to, d); err != nil {
				return err
			}

		}
	}
}

// Запись в сокет из from
func (t *Telnet) Write(ctx context.Context, from io.Reader) error {
	return t.rw(ctx, from, t.conn)

}

// Чтение из сокета в to
func (t *Telnet) Read(ctx context.Context, to io.Writer) error {
	return t.rw(ctx, t.conn, to)
}

func main() {
	var (
		timeout time.Duration
	)

	flag.DurationVar(&timeout, "timeout", defaultTimeout, "connection timeout")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())

	telnet := NewTelnet(parseAddress(), timeout)
	if err := telnet.Connect(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer telnet.Disconnect()

	// Отслеживание ошибок и отмена выполнения горутин 
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return telnet.Write(ctx, os.Stdin)
	})

	eg.Go(func() error {
		return telnet.Read(ctx, os.Stdout)
	})

	// Прерывание выполнения по сигналу
	go func() {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		<-interrupt

		cancel()
	}()
	
	if err := eg.Wait(); err != nil {
		fmt.Println(err)
	}
}

func parseAddress() string {
	host := flag.Arg(0)
	port := flag.Arg(1)

	if host == "" {
		host = defaultHost
	}

	if port == "" {
		host = defaultHost
	}

	return fmt.Sprintf("%s:%s", host, port)
}
