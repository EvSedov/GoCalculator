package routes

import (
	"github.com/evsedov/GoCalculator/web/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)

	app.Post("/getexp", handlers.CreateCalc)
}
