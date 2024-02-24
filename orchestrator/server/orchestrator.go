package orchestrator

import (
	"github.com/evsedov/GoCalculator/orchestrator/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Orchestrator struct{}

func (o *Orchestrator) Start() {
	server := fiber.New()
	server.Use(cors.New())

	routes.SetupRoutesOrchestrator(server)
	server.Listen(":8081")
}
