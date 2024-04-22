package calculateexpression

import "github.com/gofiber/fiber/v2"

func MakeHandler(s *Service) func(c *fiber.Ctx) error {
	var response *Response
	return func(c *fiber.Ctx) error {
		response = s.GetExpressionToCalculate()

		return c.Status(200).JSON(response)
	}
}
