package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	backApp := fiber.New()

	backApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Start server from backend!")
	})

	backApp.Post("/getexp", func(c *fiber.Ctx) error {
		name := c.FormValue("expression")
		return c.SendString(name)
	})

	backApp.Listen(":8080")
}
