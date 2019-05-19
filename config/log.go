package config

import (
	"log"
	"os"
)

func InitiateLog(prefix string) *log.Logger {
	return log.New(os.Stdout, prefix, log.LstdFlags|log.Lshortfile)
}
