package updatecalculated

import (
	"fmt"

	"github.com/evsedov/GoCalculator/orchestrator/entities"
	"github.com/gofiber/fiber/v2"
)

func MakeHandler(s *Service) func(c *fiber.Ctx) error {
	var response *Response
	return func(c *fiber.Ctx) error {
		var expression *entities.Expression
		if err := c.BodyParser(&expression); err != nil {
			c.Status(fiber.StatusBadRequest)
			response.Message = err.Error()

			return c.JSON(response)
		}

		response = s.UpdateCalculatedExpression(expression)
		fmt.Println("Expression:", &expression)

		return c.Status(200).JSON(response)
	}
}
