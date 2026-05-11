package main

import (
	"fmt"
	"time"
)

var (
	Debug       bool
	LogLevel    = "info"
	startUptime = time.Now()
)

func main() {
	fmt.Println(Debug, LogLevel, startUptime)
}
