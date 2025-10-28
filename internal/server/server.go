package server

import (
	"fmt"
	"log"
	"net"
	"serv-v3/internal/http1"
	"serv-v3/internal/logx"
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
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	logx.Info(2, "Server listening on %s", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("[WARM] Failed to accept connection: %v", err)
		}
		go handleConn(conn, router)
	}
}
