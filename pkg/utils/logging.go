package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func SetupLogger() (*log.Logger, error) {
	logger := log.New()
	setLoggerConfig(logger)

	logger.SetOutput(os.Stdout)
	return logger, nil
}

func setLoggerConfig(l *log.Logger) {
	l.SetFormatter(&log.JSONFormatter{})
	l.SetReportCaller(true)

	if os.Getenv("APP_ENV") == "production" {
		l.SetLevel(log.InfoLevel)
	} else {
		l.SetLevel(log.DebugLevel)
	}
}
