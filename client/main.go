package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	addr := "localhost:6001"
	requests := 1000
	start := time.Now()
	for i := 0; i < requests; i++ {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			fmt.Println("dial error", err.Error())
			continue
		}
		req := "GET / HTTP/1.1\r\nHost: localhost\r\nConnection: close\r\n\r\n"
		conn.Write([]byte(req))
		buf := make([]byte, 512)
		conn.Read(buf)
		conn.Close()
	}
	elapsed := time.Since(start)
	fmt.Printf("Отправлено %d запросов за %v\n", requests, elapsed)
	fmt.Printf("Среднее время на один запрос: %v\n", elapsed/time.Duration(requests))
}
