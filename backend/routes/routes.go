package routes

import (
	"github.com/evsedov/GoCalculator/backend/handlers"
	"github.com/evsedov/GoCalculator/backend/storage"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	expressionHandler := &handlers.ExpressionHandler{
		Storage: &storage.ExpressionStorage{
			Expressions: make(map[string]storage.Expression),
		},
	}

	app.Post("/expressions", expressionHandler.CreateExpression)

	app.Get("/expressions/:expressionId", expressionHandler.GetExpressionById)

	app.Get("/operations", func(c *fiber.Ctx) error {

		return c.SendString("Return lists of operations")
	})

	app.Get("/task", func(c *fiber.Ctx) error {

		return c.SendString("Return subexpression for calculate")
	})

	app.Post("/task/:taskId", func(c *fiber.Ctx) error {

		return c.SendString("Get result of calculate for subexpression on " + c.Params("taskId"))
	})

}
