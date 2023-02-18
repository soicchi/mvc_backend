package main

import (
	"github.com/soicchi/chatapp_backend/pkg/database"
)

func main() {
	err := database.SetupDB()
	if err != nil {
		panic(err)
	}
}