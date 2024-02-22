package routes

import (
	"github.com/evsedov/GoCalculator/calculator/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutesCalculator(app *fiber.App) {
	calculatorHandlers := &handlers.CalculatorHendler{}

	app.Get("/start", calculatorHandlers.Start)
}
