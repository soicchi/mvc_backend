package logging

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func SetupLogger() (*log.Logger, error) {
	logger := log.New()
	logger.SetFormatter(&log.JSONFormatter{})
	logger.SetReportCaller(true)

	if os.Getenv("ENV") == "production" {
		logger.SetLevel(log.InfoLevel)
	} else {
		logger.SetLevel(log.DebugLevel)
	}

	file, err := getLogFile()
	if err != nil {
		return nil, err
	}

	logger.SetOutput(file)
	return logger, nil
}

func getLogFile() (*os.File, error) {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	return file, nil
}
