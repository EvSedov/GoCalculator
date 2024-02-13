package handlers

import "github.com/gofiber/fiber/v2"

var body []byte

func HomeHandler(c *fiber.Ctx) error {

	return c.SendString("Start server from backend!")
}

func GetExpressions(c *fiber.Ctx) error {

	return c.SendString(string(body))
}

func SetExpression(c *fiber.Ctx) error {
	body = c.Body()

	return c.SendStatus(200)
}
