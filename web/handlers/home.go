package handlers

import "github.com/gofiber/fiber/v2"

type Calc struct {
	Calc string `json:"calc"`
}

func Home(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Distributed calculator",
	})
}
