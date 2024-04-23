package routes

import (
	"github.com/evsedov/GoCalculator/calculator/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutesCalculator(app *fiber.App) {
	handlers := &handlers.CalculatorHendler{}

	app.Get("/start", handlers.Start)
}
