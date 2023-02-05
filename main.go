package main

import (
	"github.com/MohammedKamle/banking/app"
	"github.com/MohammedKamle/banking/logger"
)

func main() {
	logger.Info("Starting the application....")
	app.StartApplication()
}
