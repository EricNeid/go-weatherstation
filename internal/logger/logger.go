package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	internalLog *log.Logger
)

// Log provides logging methods. Provide the given context (file or package name for example)
// to be displayed in the log messages.
type Log struct {
	Context string
}

// Init configures logger to write to file
func Init() {
	logPath := "go-weatherstation.log"
	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	internalLog = log.New(f, "", log.LstdFlags|log.Lshortfile)
	internalLog.Println("LogFile : " + logPath)
}

// D writes the given message to log output
func (l *Log) D(method string, msg string) {
	internalLog.Printf("%s: %s: %s\n", l.Context, method, msg)
}

// E writes the given error to log output
func (l *Log) E(method string, err error) {
	internalLog.Printf(fmt.Sprintf("ERROR: %s: %s: %s\n", l.Context, method, err.Error()), err)
}
