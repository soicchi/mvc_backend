package main

import (
	"github.com/soicchi/chatapp_backend/pkg/database"
	"github.com/soicchi/chatapp_backend/pkg/logging"
	"github.com/soicchi/chatapp_backend/pkg/router"
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

	router.Run()
}