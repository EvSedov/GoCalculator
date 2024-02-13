package main

import (
	"github.com/evsedov/GoCalculator/backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	backApp := fiber.New()
	backApp.Use(cors.New())

	routes.SetupRoutes(backApp)

	backApp.Listen(":8081")
}
