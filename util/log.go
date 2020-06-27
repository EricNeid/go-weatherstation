package util

import "log"

// Log provides logging methods. Provide the given context (file or package name for example)
// to be displayed in the log messages.
type Log struct {
	Context string
}

// D writes the given message to log output
func (l *Log) D(method string, msg string) {
	log.Printf("%s: %s: %s\n", l.Context, method, msg)
}

// E writes the given error to log output
func (l *Log) E(method string, err error) {
	log.Printf("ERROR: %s: %s: %s\n", l.Context, method, err.Error())
}
