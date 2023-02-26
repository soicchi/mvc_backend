package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func SetupLogger(logFilePath string) (*log.Logger, error) {
	logger := log.New()
	setLoggerConfig(logger)

	file, err := getLogFile(logFilePath)
	if err != nil {
		return nil, err
	}

	logger.SetOutput(file)
	return logger, nil
}

func getLogFile(logFilePath string) (*os.File, error) {
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	return file, nil
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
