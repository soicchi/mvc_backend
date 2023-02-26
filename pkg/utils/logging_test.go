package utils

import (
	"os"
	"testing"
)

func TestSetupLogger(t *testing.T) {
	testLogFilePath := "test.log"
	_, err := SetupLogger(testLogFilePath)
	if err != nil {
		t.Errorf("Error setting up logger: %v", err)
	}
	defer os.Remove(testLogFilePath)
}

func TestGetLogFile(t *testing.T) {
	testLogFilePath := "test.log"
	_, err := getLogFile(testLogFilePath)
	if err != nil {
		t.Errorf("Error getting log file: %v", err)
	}
	defer os.Remove(testLogFilePath)
}
