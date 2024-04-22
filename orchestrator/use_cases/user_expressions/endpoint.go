package userexpressions

import "github.com/gofiber/fiber/v2"

func MakeHandler(s *Service) func(c *fiber.Ctx) error {
	var response *Response
	return func(c *fiber.Ctx) error {
		email := c.Locals("email").(string)
		response = s.GetUserExpressions(email)

		return c.Status(200).JSON(response)
	}
}
