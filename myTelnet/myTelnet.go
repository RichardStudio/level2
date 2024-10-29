package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	// Парсинг аргументов командной строки
	timeout := flag.Duration("timeout", 10*time.Second, "connection timeout")
	flag.Parse()

	if flag.NArg() < 2 {
		fmt.Println("Usage: go-telnet [--timeout=10s] host port")
		return
	}

	host := flag.Arg(0)
	port := flag.Arg(1)

	address := net.JoinHostPort(host, port)

	// Подключение к серверу
	conn, err := net.DialTimeout("tcp", address, *timeout)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to", address)

	// Канал для сигналов завершения
	done := make(chan struct{})

	// Чтение данных из сокета и вывод в STDOUT
	go func() {
		ioCopy(os.Stdout, conn)
		fmt.Println("Connection closed by server")
		done <- struct{}{}
	}()

	// Чтение данных из STDIN и запись в сокет
	go func() {
		ioCopy(conn, os.Stdin)
		done <- struct{}{}
	}()

	// Ожидание завершения
	<-done
}

func ioCopy(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
