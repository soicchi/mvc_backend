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
