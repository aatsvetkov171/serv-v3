package logx

import (
	"fmt"
	"log"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorYellow = "\033[33m"
	ColorCyan   = "\033[36m"
)

var logger = log.Default()

func Info(depth int, msg string, args ...interface{}) {
	logger.Output(depth, ColorCyan+fmt.Sprintf("[INFO] "+msg, args...)+ColorReset)
}
