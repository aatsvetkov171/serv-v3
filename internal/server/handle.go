package server

import (
	"fmt"
	"net"
	"serv-v3/internal/http1"
	"time"
)

func handleConn(conn net.Conn) {
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
		request := http1.NewRequest(&FLineB)
		request.GetConnHeaders(conn, reader)
		request.ReadBody(conn, reader)
		fmt.Println("ТЕЛО", request.GetBody())

		if request.GetHeaders()["connection"] == "close" {
			keepAlive = false
		}
		//response_content := "<h1>Hello v3</h1>"
		response := http1.NewResponse(200, "<h1>Hello 876612412</h1>")
		response.AddHeaders(http1.DefaultHeaders)
		response.Write(conn)
		if CountMessage >= maxCountMessage {
			fmt.Println("Достигнут лимит кол сообщений")
			return
		}
		fmt.Println("Время одного запроса:", time.Since(start))
	}

}
