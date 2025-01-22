package logmanager

import (
	"log"
	"os"
)

var (
	// InfoLogger logs informational messages to stdout
	InfoLogger *log.Logger
	// ErrorLogger logs error messages to stderr
	ErrorLogger *log.Logger
)

// init initializes the loggers for Info and Error automatically
func init() {
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Info logs informational messages to stdout
func Info(msg string) {
	InfoLogger.Println(msg)
}

// Infof logs formatted informational messages to stdout
func Infof(format string, v ...interface{}) {
	InfoLogger.Printf(format, v...)
}

// Error logs error messages to stderr
func Error(err error) {
	if err != nil {
		ErrorLogger.Println(err)
	}
}

// Errorf logs formatted error messages to stderr
func Errorf(format string, v ...interface{}) {
	ErrorLogger.Printf(format, v...)
}
