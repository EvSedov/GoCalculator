package calculator

import (
	"github.com/evsedov/GoCalculator/calculator/routes"
	"github.com/gofiber/fiber/v2"
	// "github.com/google/uuid"
)

type Calculator struct {
	ServerId string `json:"server_id"`
}

func (o *Calculator) Start() {
	// o.ServerId = uuid.New().String()
	serverCalculator := fiber.New()

	routes.SetupRoutesCalculator(serverCalculator)
	serverCalculator.Listen(":85")
}
