package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type (
	MainHandler struct{}
)

func (mh *MainHandler) GetMain(c *fiber.Ctx) error {

	return c.Render("index", fiber.Map{
		"Title": "Distributed calculator",
	})
}
