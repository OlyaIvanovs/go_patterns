package main

import (
	"fmt"
	"sync"
)

// MyLogger is the struct we want to make singleton
type MyLogger struct {
	logLevel int
}

func (l *MyLogger) Log(s string) {
	fmt.Println(l.logLevel, ":", s)
}

func (l *MyLogger) SetLogLevel(level int) {
	l.logLevel = level
}

// the logger Instance
var logger *MyLogger

var once sync.Once

//Logger class instance
func getLoggerInstance() *MyLogger {
	once.Do(func() {
		fmt.Println("Creating Logger Instance")
		logger = &MyLogger{}
	})
	fmt.Println("Returning Logger Instance")
	return logger
}
