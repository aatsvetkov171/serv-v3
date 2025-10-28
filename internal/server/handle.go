package server

import (
	"fmt"
	"net"
	"serv-v3/internal/http1"
	"time"
)

func handleConn(conn net.Conn, router *http1.Router) {
	defer func() {
		fmt.Println("close conn", conn.RemoteAddr())
		conn.Close()
	}()
	keepAlive := true
	reader := CreateReader(conn)
	CountMessage := 0
	for {
		start := time.Now()
		if !keepAlive {
			break
		}
		CountMessage += 1
		FLineB, err := ReadFirstLine(conn, reader)
		if err != nil {
			if nettErr, ok := err.(net.Error); ok && nettErr.Timeout() {
				fmt.Println("TIME OUT")
				return
			}
			fmt.Println("getting first line error:", err.Error())
			return
		}
		if !IsHttp(&FLineB) {
			fmt.Println("Server cant working with non http proto..")
			return
		}
		//----------------REQUEST
		request := http1.NewRequest(&FLineB)
		request.GetConnHeaders(conn, reader)
		request.ReadBody(conn, reader)
		//----------------RESPONSE
		responseFunc, ok := router.FindHandler(request.GetMethod(), request.GetPath())
		if !ok {
			return
		}

		response := responseFunc(&request)
		response.Write(conn)

		if request.GetHeaders()["connection"] == "close" {
			keepAlive = false
		}
		//response_content := "<h1>Hello v3</h1>
		if CountMessage >= maxCountMessage {
			fmt.Println("Достигнут лимит кол сообщений")
			return
		}
		fmt.Println("Время одного запроса:", time.Since(start))
	}

}
