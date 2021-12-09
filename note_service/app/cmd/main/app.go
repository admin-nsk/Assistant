package main

import (
	"note_service/app/internal/router"
	"note_service/app/pkg/logging"
)

func main() {
	logging.Init()
	logger := logging.Getlogger()
	logger.Println("init logger")

	defer router.Init()

	logger.Println("Application inicialized and started")
}
