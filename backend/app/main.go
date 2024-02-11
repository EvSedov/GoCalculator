package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	var body []byte
	backApp := fiber.New()
	backApp.Use(cors.New())

	backApp.Get("/", func(c *fiber.Ctx) error {

		return c.SendString("Start server from backend!")
	})

	backApp.Get("/getexp", func(c *fiber.Ctx) error {

		return c.SendString(string(body))
	})

	backApp.Post("/getexp", func(c *fiber.Ctx) error {
		body = c.Body()

		return c.SendStatus(200)
	})

	backApp.Listen(":8081")
}
