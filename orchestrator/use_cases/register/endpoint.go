package register

import (
	"github.com/evsedov/GoCalculator/orchestrator/entities"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
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

		password, _ := bcrypt.GenerateFromPassword([]byte(payload["password"]), 14)
		user := entities.User{
			Name:     payload["name"],
			Email:    payload["email"],
			Password: password,
		}
		response := s.CreateUser(user)
		c.Status(fiber.StatusOK)

		return c.JSON(response)
	}
}
