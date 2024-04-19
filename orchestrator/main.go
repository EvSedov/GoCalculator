package main

import (
	"github.com/evsedov/GoCalculator/orchestrator/app"
)

func main() {
	orchestrator := &app.Orchestrator{}
	orchestrator.Run()
}
