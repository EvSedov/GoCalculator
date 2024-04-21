package register

import (
	"fmt"

	"github.com/evsedov/GoCalculator/orchestrator/entities"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func MakeHandler(s *Service) func(c *fiber.Ctx) error {
	var response *Response
	return func(c *fiber.Ctx) error {
		var payload map[string]string
		if err := c.BodyParser(&payload); err != nil {
			c.Status(fiber.StatusBadRequest)
			response.Error = err.Error()

			return c.JSON(response)
		}

		password, _ := bcrypt.GenerateFromPassword([]byte(payload["password"]), 14)
		user := entities.User{
			Name:     payload["name"],
			Email:    payload["email"],
			Password: password,
		}
		response = s.CreateUser(user)
		if response.Error != "" {
			c.Status(fiber.StatusForbidden)

			return c.JSON(response)
		}

		token := response.Token
		c.Set("Access-Control-Expose-Headers", "Authorization")
		c.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		c.Status(fiber.StatusCreated)

		return c.JSON(response)
	}
}
