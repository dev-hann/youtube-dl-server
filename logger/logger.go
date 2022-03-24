package logger

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func InitLogger() *log.Entry {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	return log.WithFields(
		log.Fields{},
	)
}
