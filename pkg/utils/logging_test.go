package utils

import (
	"testing"
)

func TestSetupLogger(t *testing.T) {
	_, err := SetupLogger()
	if err != nil {
		t.Errorf("Error setting up logger: %v", err)
	}
}
