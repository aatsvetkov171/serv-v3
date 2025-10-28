package http1

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
)

type Request struct {
	method  string
	url     string
	proto   string
	headers map[string]string
	body    string
}

func NewRequest(FLineB *[]byte) Request {
	slice := strings.Fields(string(*FLineB))
	newRequest := Request{
		method:  slice[0],
		url:     slice[1],
		proto:   slice[2],
		headers: make(map[string]string),
		body:    "",
	}
	return newRequest
}

func (r *Request) GetConnHeaders(conn net.Conn, reader *bufio.Reader) {
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("reading headers error:", err.Error())
			break
		}
		lineS := string(line)
		fmt.Println(lineS)
		if lineS == "\r\n" {
			break
		}
		slice := strings.SplitN(lineS, ":", 2)
		if len(slice) < 2 {
			fmt.Println("Invalid header slice")
			continue
		}
		key := strings.ToLower(strings.TrimSpace(slice[0]))
		val := strings.TrimSpace(slice[1])
		r.headers[key] = val
	}
}

func (r *Request) GetHeaders() map[string]string {
	return r.headers
}

func (r *Request) ReadBody(conn net.Conn, reader *bufio.Reader) {
	if r.method == "POST" || r.method == "PUT" || r.method == "PATCH" {
		bufSize, _ := strconv.Atoi(r.headers["content-length"])
		buff := make([]byte, bufSize)
		reader.Read(buff)
		r.body = string(buff)
	}
}

func (r *Request) GetBody() string {
	return r.body
}
