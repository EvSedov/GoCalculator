package routes

import (
	// "github.com/evsedov/GoCalculator/orchestrator/handlers"
	"github.com/evsedov/GoCalculator/orchestrator/middleware"
	"github.com/evsedov/GoCalculator/orchestrator/storage"
	"github.com/evsedov/GoCalculator/orchestrator/use_cases/calculate_expression"
	"github.com/evsedov/GoCalculator/orchestrator/use_cases/create_expression"
	"github.com/evsedov/GoCalculator/orchestrator/use_cases/login"
	"github.com/evsedov/GoCalculator/orchestrator/use_cases/register"
	updatecalculated "github.com/evsedov/GoCalculator/orchestrator/use_cases/update_calculated"
	"github.com/evsedov/GoCalculator/orchestrator/use_cases/user_expressions"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	storage := storage.NewStorage()
	// orchestratorHandler := &handlers.Handler{}

	registerService := register.NewService(storage)
	loginService := login.NewService(storage)
	createExpressionService := createexpression.NewService(storage)
	getUserExpressionService := userexpressions.NewService(storage)
	getExpressionToCalculateService := calculateexpression.NewService(storage)
	updateCalculatedExpressionService := updatecalculated.NewService(storage)

	api := app.Group("api")
	v1 := api.Group("v1")
	v1.Post("/register", register.MakeHandler(registerService))
	v1.Post("/login", login.MakeHandler(loginService))

	v1.Use("/expressions", middleware.AuthenticateMiddleware)
	v1.Post("/expressions", createexpression.MakeHandler(createExpressionService))
	v1.Get("/expressions", userexpressions.MakeHandler(getUserExpressionService))

	v1.Get("/calculate", calculateexpression.MakeHandler(getExpressionToCalculateService))
	v1.Post("/calculate", updatecalculated.MakeHandler(updateCalculatedExpressionService))

	// app.Get("/expressions/:expression_id", orchestratorHandler.GetExpressionById)
	// app.Get("/operations", func(c *fiber.Ctx) error {
	// 	return c.SendString("Return lists of operations")
	// })
}
