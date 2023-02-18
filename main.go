package main

import (
	"github.com/soicchi/chatapp_backend/pkg/database"
	"github.com/soicchi/chatapp_backend/pkg/logging"
)

func main() {
	logger, err := logging.SetupLogger()
	if err != nil {
		panic(err)
	}

	err = database.SetupDB()
	if err != nil {
		logger.Fatal(err)
	}
}