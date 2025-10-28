package server

import (
	"fmt"
	"net"
)

func GoServer() {
	listener, err := net.Listen(network, addr)
	if err != nil {
		fmt.Println("create listener error:", err.Error())
	}
	defer listener.Close()
	//router := http1.NewRouter()
	fmt.Println("Server listening:", addr, "....")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("connection error:", err.Error())
		}
		go handleConn(conn)
	}
}
