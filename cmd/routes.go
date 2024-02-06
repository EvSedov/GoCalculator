package main

import (
	"github.com/EvSedov/GoCalculator/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)

	app.Post("/calc", handlers.CreateCalc)
}
