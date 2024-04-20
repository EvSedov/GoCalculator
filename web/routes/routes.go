package routes

import (
	"github.com/evsedov/GoCalculator/web/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	mainHandler := &handlers.MainHandler{}

	app.Get("/", mainHandler.GetHome)
	app.Get("/login", mainHandler.GetLogin)
	app.Get("/register", mainHandler.GetRegister)
}
