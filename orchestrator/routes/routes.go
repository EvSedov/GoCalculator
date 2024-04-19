package routes

import (
	// "github.com/evsedov/GoCalculator/orchestrator/handlers"
	"github.com/evsedov/GoCalculator/orchestrator/storage"
	"github.com/evsedov/GoCalculator/orchestrator/use_cases/login"
	"github.com/evsedov/GoCalculator/orchestrator/use_cases/register"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	storage := storage.NewStorage()
	// orchestratorHandler := &handlers.Handler{}

	registerService := register.NewService(storage)
	loginService := login.NewService(storage)

	app.Post("/api/register", register.MakeHandler(registerService))
	app.Post("/api/login", login.MakeHandler(loginService))

	// app.Get("/expressions", orchestratorHandler.GetExpressions)

	// app.Post("/expressions", orchestratorHandler.CreateExpression)

	// app.Get("/expressions/:expression_id", orchestratorHandler.GetExpressionById)

	// app.Get("/operations", func(c *fiber.Ctx) error {

	// 	return c.SendString("Return lists of operations")
	// })

	// app.Get("/task", orchestratorHandler.GetValidExpressionToWork)

	// app.Post("/task", orchestratorHandler.UpdateExpressionInWork)

}
