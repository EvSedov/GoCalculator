package routes

import (
	"github.com/evsedov/GoCalculator/orchestrator/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutesOrchestrator(app *fiber.App) {
	orchestratorHandler := &handlers.OrchestratorHandler{}

	app.Get("/expressions", orchestratorHandler.GetExpressions)

	app.Post("/expressions", orchestratorHandler.CreateExpression)

	app.Get("/expressions/:expression_id", orchestratorHandler.GetExpressionById)

	app.Get("/operations", func(c *fiber.Ctx) error {

		return c.SendString("Return lists of operations")
	})

	app.Get("/task", orchestratorHandler.GetValidExpressionToWork)

	app.Post("/task", orchestratorHandler.UpdateExpressionInWork)

}
