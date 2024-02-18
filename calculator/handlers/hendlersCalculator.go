package handlers

import (
	"github.com/gofiber/fiber/v2"
	// "github.com/google/uuid"
	// "github.com/evsedov/GoCalculator/backend/storage"
	// "github.com/evsedov/GoCalculator/backend/utils"
)

type (
	CalculatorHendler struct {
	}

	CreateRegistrationRespons struct {
		ServerID   string `form:"request_id"`
		Expression string `form:"expression"`
	}
)

func (calc *CalculatorHendler) Registration(ctx *fiber.Ctx) error {
	return ctx.SendString("Server calculator is started!")
}
