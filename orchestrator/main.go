package main

import (
	"github.com/evsedov/GoCalculator/orchestrator/server"
)

func main() {
	orchestrator := &orchestrator.Orchestrator{}

	orchestrator.Start()
}
