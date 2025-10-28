package server

import (
	"fmt"
	"net"
	"serv-v3/internal/http1"
)

func GoServer() {
	listener, err := net.Listen(network, addr)
	if err != nil {
		fmt.Println("create listener error:", err.Error())
	}
	defer listener.Close()
	router := http1.NewRouter()
	router.Handle("GET", "/hello", http1.PathHello)
	router.Handle("GET", "/about", http1.PathAbout)
	fmt.Println("Server listening:", addr, "....")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("connection error:", err.Error())
		}
		go handleConn(conn, router)
	}
}
