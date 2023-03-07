package utils

import (
	"os"
	"io"

	log "github.com/sirupsen/logrus"
)

func SetupLogger(logFileName string) (*log.Logger, error) {
	logger := log.New()
	setLoggerConfig(logger)

	file, err := getLogFile(logFileName)
	if err != nil {
		return nil, err
	}

	multiOutput := io.MultiWriter(file, os.Stdout)
	logger.SetOutput(multiOutput)
	return logger, nil
}

func getLogFile(logFileName string) (*os.File, error) {
	var prefix string
	if os.Getenv("APP_ENV") == "production" {
		prefix = "production"
	} else {
		prefix = "development"
	}

	logFilePath := "logs/" + prefix + "/" + logFileName
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
