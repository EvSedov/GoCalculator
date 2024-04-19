package app

import (
	"github.com/evsedov/GoCalculator/orchestrator/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Orchestrator struct{}

func (o *Orchestrator) Run() {
	server := fiber.New()
	server.Use(cors.New())

	routes.Setup(server)
	server.Listen(":8081")
}
