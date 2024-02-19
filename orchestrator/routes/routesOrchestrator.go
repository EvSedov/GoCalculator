package routes

import (
	"github.com/evsedov/GoCalculator/orchestrator/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutesOrchestrator(app *fiber.App) {
	orchestratorHandler := &handlers.OrchestratorHandler{}

	app.Get("/expressions", orchestratorHandler.GetExpressions)

	app.Post("/expressions", orchestratorHandler.CreateExpression)

	app.Get("/expressions/:expressionId", orchestratorHandler.GetExpressionById)

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
