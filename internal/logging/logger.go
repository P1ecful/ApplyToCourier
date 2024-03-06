package logging

import (
	"log"
	"os"
)

func NewLogger() *log.Logger {
	return log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
}
