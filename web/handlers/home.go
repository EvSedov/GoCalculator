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

func CreateCalc(c *fiber.Ctx) error {
	calc := Calc{}
	if err := c.BodyParser(&calc); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(calc)
}
