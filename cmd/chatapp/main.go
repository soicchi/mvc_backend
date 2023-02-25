package main

import (
	"github.com/soicchi/chatapp_backend/pkg/database"
	"github.com/soicchi/chatapp_backend/pkg/router"
	"github.com/soicchi/chatapp_backend/pkg/utils"
)

func main() {
	logger, err := utils.SetupLogger()
	if err != nil {
		panic(err)
	}

	err = database.SetupDB()
	if err != nil {
		logger.Fatal(err)
	}

	router.Run()
}
