package createexpression

import (
	"github.com/evsedov/GoCalculator/orchestrator/entities"
	"github.com/evsedov/GoCalculator/orchestrator/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func MakeHandler(s *Service) func(c *fiber.Ctx) error {
	var response *Response
	return func(c *fiber.Ctx) error {
		var payload map[string]string
		if err := c.BodyParser(&payload); err != nil {
			response.Message = err.Error()
			c.Status(fiber.StatusBadRequest)

			return c.JSON(response)
		}

		var expression entities.Expression
		expression.Email = c.Locals("email").(string)
		expression.ExpressionId = uuid.New().String()
		expression.RequestID = payload["expression_id"]
		expression.RequestID = payload["request_id"]
		expression.Expression = utils.DelSpaceFromString((payload["expression"]))

		if utils.IsValidExpression(expression.Expression) {
			expression.Message = "the expression will be calculated soon"
			expression.State = "valid"
		} else {
			expression.Message = "expression parsing error"
			expression.State = "error"
		}

		response = s.CreateExpression(expression)

		return c.Status(200).JSON(expression)

	}
}
