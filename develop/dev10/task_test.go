package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func setupTestserver(addr string) (net.Listener, error) {
	conn, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			c, _ := conn.Accept()
			defer c.Close()

			go func() {
				r := bufio.NewReader(c)
				for {
					d, err := r.ReadString('\n')
					fmt.Println("readed", d)
					if err != nil {
						return
					}

					fmt.Fprintf(c, d)
				}
			}()
		}
	}()

	return conn, err
}

func TestTelnet(t *testing.T) {
	addr := "127.0.0.1:5055"
	timeout := 5 * time.Second
	ctx := context.Background()

	expected := "Hello test\n"
	from := bytes.NewBufferString(expected)
	to := bytes.NewBufferString("")

	conn, err := setupTestserver(addr)
	assert.NoError(t, err)

	telnet := NewTelnet(conn.Addr().String(), timeout)
	err = telnet.Connect()
	assert.NoError(t, err)

	go telnet.Write(ctx, from)
	go telnet.Read(ctx, to)

	time.Sleep(1 * time.Second)

	assert.Equal(t, expected, to.String())
}
