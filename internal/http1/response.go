package http1

import (
	"net"
	"strconv"
)

var status = map[int]string{
	200: "OK",
	404: "Not Found",
}

var DefaultHeaders = map[string]string{
	"content-type": "text/html; charset=utf-8",
	"connection":   "keep-alive",
}

type Response struct {
	proto      string
	statusCode int
	statusMess string
	headers    map[string]string
	body       string
}

type Response404 struct {
	statusCode int
	statusMess string
	body       string
}

func NewResponse404() *Response404 {
	return &Response404{
		statusCode: 404,
		statusMess: "Not found",
		body:       "Not Found",
	}
}

func NewResponse(statusCode int, body string) *Response {
	newResponse := Response{
		proto:      "HTTP/1.1",
		statusCode: statusCode,
		statusMess: status[statusCode],
		headers:    make(map[string]string),
		body:       body,
	}
	return &newResponse
}

func (r *Response) AddHeaders(h map[string]string) {
	for key, val := range h {
		r.headers[key] = val
	}
	r.headers["content-length"] = strconv.Itoa(len(r.body))
}

func (r *Response) Write(conn net.Conn) {
	httpText := r.proto + " " + strconv.Itoa(r.statusCode) + " " + r.statusMess + "\r\n"
	for k, v := range r.headers {
		httpText += k + ":" + v + "\r\n"
	}
	httpText += "\r\n" + r.body
	conn.Write([]byte(httpText))
}

func (r *Response404) Write(conn net.Conn) {
	httpText := "HTTP/1.1 " + strconv.Itoa(r.statusCode) + " " + r.statusMess + "\r\n"
	httpText += "\r\n" + r.body
	conn.Write([]byte(httpText))
}
