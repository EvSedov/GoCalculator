package routes

import (
	"github.com/evsedov/GoCalculator/web/handlers"
	"github.com/evsedov/GoCalculator/web/storage"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	mainHandler := &handlers.MainHandler{
		Storage: &storage.WebExpressionStorage{
			Expressions: make([]storage.Expression, 0),
		},
	}

	app.Get("/", mainHandler.GetMain)
	app.Get("/expressions", mainHandler.GetExpressions)
}
