package main

import (
	"github.com/evsedov/GoCalculator/orchestrator/database"
	"github.com/evsedov/GoCalculator/orchestrator/server"
)

func main() {
	database.ConnectDb()

	orchestrator := &orchestrator.Orchestrator{}

	orchestrator.Start()
}
