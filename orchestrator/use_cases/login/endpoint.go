package login

import (
	"github.com/evsedov/GoCalculator/orchestrator/entities"
	"github.com/gofiber/fiber/v2"
)

func MakeHandler(s *Service) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var payload map[string]string
		if err := c.BodyParser(&payload); err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err,
			})
		}

		user := entities.User{
			Email:    payload["email"],
			Password: []byte(payload["password"]),
		}
		response := s.Login(user)
		c.Status(fiber.StatusOK)

		return c.JSON(response)
	}
}
