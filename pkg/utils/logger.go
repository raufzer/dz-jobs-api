package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func InitLogger() {

	logLevel, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = log.DebugLevel
	}

	log.SetLevel(logLevel)
	log.SetFormatter(&log.JSONFormatter{})
}
