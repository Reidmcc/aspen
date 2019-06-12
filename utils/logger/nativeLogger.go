package logger

import (
	"io"
	"log"
)

// nativeLogger is a Go native logger
type nativeLogger struct {
}

// Info impl
func (l *nativeLogger) Info(msg string) {
	log.Println(msg)
}

// Infof impl
func (l *nativeLogger) Infof(msg string, args ...interface{}) {
	log.Printf(msg, args...)
}

// Error impl
func (l *nativeLogger) Error(msg string) {
	log.Print(msg)
}

// Efforf impl
func (l *nativeLogger) Errorf(msg string, args ...interface{}) {
	log.Printf(msg, args...)
}

// SetOutput impl
func (l *nativeLogger) SetOutput(w io.Writer) {
	log.SetOutput(w)
}

// ensure it implements Logger
var _ Logger = &nativeLogger{}

// MakenativeLogger is the factory method
func MakenativeLogger() Logger {
	return &nativeLogger{}
}
