package main

import (
	"github.com/evsedov/GoCalculator/orchestrator/app"
	"github.com/evsedov/GoCalculator/orchestrator/database"
)

func main() {
	database.Connect()

	orchestrator := &app.Orchestrator{}
	orchestrator.Run()
}
