package logger

import (
	"log"
	"os"
)

// Different types of loggers to log into os.Stderr
var (
	InfoLogger    = log.New(os.Stderr, "INFO: ", log.Ltime|log.Lshortfile)
	WarningLogger = log.New(os.Stderr, "WARNING: ", log.Ltime|log.Lshortfile)
	ErrorLogger   = log.New(os.Stderr, "ERROR: ", log.Ltime|log.Lshortfile)
)
