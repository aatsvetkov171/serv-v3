package server

import (
	"bufio"
	"net"
	"strings"
	"time"
)

func isBlank(FlineB []byte) bool {
	return len(FlineB) == 0
}

func ReadFirstLine(conn net.Conn, reader *bufio.Reader) ([]byte, error) {
	conn.SetReadDeadline(time.Now().Add(time.Duration(keepAliveMaxSeconds) * time.Second))
	FlineB, err := reader.ReadBytes('\n')
	if err != nil {
		return FlineB, err
	}
	blank := isBlank(FlineB)
	if blank {
		return FlineB, err
	}
	return FlineB, err
}

func CreateReader(conn net.Conn) *bufio.Reader {
	reader := bufio.NewReader(conn)
	return reader
}

func IsHttp(FLineB *[]byte) bool {
	FLineS := string(*FLineB)
	FLineSSlice := strings.Fields(FLineS)
	return ((len(FLineSSlice) == 3) && (FLineSSlice[2] == "HTTP/1.1"))
}
