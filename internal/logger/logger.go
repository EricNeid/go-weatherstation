package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

var internalLog = log.New(os.Stderr, "", log.LstdFlags)

// Log provides logging methods. Provide the given context (file or package name for example)
// to be displayed in the log messages.
type Log struct {
	Context string
}

// Init configures logger to write to file
func Init() {
	logPath := "go-weatherstation.log"

	// create logfile
	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0o666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	// create log writer, writes to file and console
	logWriter := io.MultiWriter(os.Stdout, f)

	internalLog = log.New(logWriter, "", log.LstdFlags|log.Lshortfile)
	internalLog.Println("LogFile : " + logPath)
}

// D writes the given message to log output
func (l *Log) D(method, msg string) {
	internalLog.Printf("%s: %s\n", method, msg)
}

// E writes the given error to log output
func (l *Log) E(method string, err error) {
	internalLog.Printf(fmt.Sprintf("ERROR: %s: %s\n", method, err.Error()), err)
}
