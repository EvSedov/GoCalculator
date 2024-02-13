package routes

import (
	"github.com/evsedov/GoCalculator/backend/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(backApp *fiber.App) {
	backApp.Get("/", handlers.HomeHandler)

	backApp.Get("/expressions", handlers.GetExpressions)

	backApp.Post("/expressions", handlers.SetExpression)

	backApp.Get("/expressions/:expressionId", func(c *fiber.Ctx) error {

		return c.SendString("Get expression id = " + c.Params("expressionId"))
	})

	backApp.Get("/operations", func(c *fiber.Ctx) error {

		return c.SendString("Return lists of operations")
	})

	backApp.Get("/task", func(c *fiber.Ctx) error {

		return c.SendString("Return subexpression for calculate")
	})

	backApp.Post("/task/:taskId", func(c *fiber.Ctx) error {

		return c.SendString("Get result of calculate for subexpression on " + c.Params("taskId"))
	})

}
